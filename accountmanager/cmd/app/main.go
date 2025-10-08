package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github/com/CargoMan0/GoPay/accountmanager/internal/config"
	"github/com/CargoMan0/GoPay/accountmanager/internal/eventsender"
	"github/com/CargoMan0/GoPay/accountmanager/internal/hasher"
	"github/com/CargoMan0/GoPay/accountmanager/internal/jwt"
	"github/com/CargoMan0/GoPay/accountmanager/internal/repository"
	"github/com/CargoMan0/GoPay/accountmanager/internal/server"
	"github/com/CargoMan0/GoPay/accountmanager/internal/service"
	"github/com/CargoMan0/GoPay/pkg/database"
	"github/com/CargoMan0/GoPay/pkg/kafka"
	"log"
	"log/slog"
	"os/signal"
	"sync"
	"syscall"
)

func main() {
	err := run()
	if err != nil {
		log.Fatalf("run() returned error: %v", err)
	}
}

func run() (err error) {
	ctx, cancel := signal.NotifyContext(
		context.Background(),
		syscall.SIGINT,
		syscall.SIGTERM,
	)
	defer cancel()

	slog.Info("Loading config...")
	cfg, err := config.Load()
	if err != nil {
		return fmt.Errorf("load config: %v", err)
	}
	slog.Info("Config loaded")

	slog.Info("Creating new database...")
	db, err := newSqlDB()
	if err != nil {
		return fmt.Errorf("new sql db: %v", err)
	}
	slog.Info("Database created")
	defer func() {
		slog.Info("Closing database connection...")
		closeErr := db.Close()
		if closeErr != nil {
			err = errors.Join(err, fmt.Errorf("close db: %v", closeErr))
		}
		slog.Info("Database connection closed")
	}()

	producer := newKafkaProducer(cfg.Producer)
	go func() {
		<-ctx.Done()
		closeErr := producer.Close()
		if closeErr != nil {
			err = errors.Join(err, fmt.Errorf("close kafka producer: %v", closeErr))
		}
	}()

	// Repository
	repo := repository.New(db)

	// Utils
	passwordHasher := hasher.NewPasswordHasher()

	// Service
	accountService := service.NewAccountService(
		repo,
		passwordHasher,
	)

	// Server
	srv := server.NewGRPC(cfg.GRPCServer, accountService)

	wg := &sync.WaitGroup{}
	wg.Add(1)

	eventSender := eventsender.New(
		cfg.EventSender,
		producer,
	)

	wg.Add(1)
	go func(ctx context.Context) {
		defer wg.Done()
		slog.Info("Starting EventSender...")
		eventSender.RunEventProcessing(ctx)

		slog.Info("EventSender stopped")
	}(ctx)

	go func() {
		defer wg.Done()
		slog.Info("Starting gRPC server...")
		srvErr := srv.Start()
		if srvErr != nil {
			err = errors.Join(err, fmt.Errorf("start gRPC server: %v", srvErr))
			// Cancel context if failed to start server
			cancel()
			return
		}

		slog.Info("gRPC server started")
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		<-ctx.Done()

		slog.Info("Shutting down gRPC server...")
		srv.Stop()
		slog.Info("gRPC server shut down")
	}()

	wg.Wait()
	return nil
}

func newSqlDB() (*sql.DB, error) {
	c := &database.Config{
		// TODO: write
	}

	return database.New(c)
}

func newKafkaProducer(cfg config.Producer) *kafka.Producer {
	c := kafka.ProducerConfig{
		// TODO: write
	}

	return kafka.NewProducer(c)
}
