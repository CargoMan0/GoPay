package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/CargoMan0/GoPay/gateway/internal/clients/grpc"
	"github.com/CargoMan0/GoPay/gateway/internal/config"
	"github.com/CargoMan0/GoPay/gateway/internal/controller"
	"github.com/CargoMan0/GoPay/gateway/internal/controller/http"
	"github.com/CargoMan0/GoPay/gateway/internal/server"
	"github.com/CargoMan0/GoPay/gateway/internal/validator"
	"github.com/CargoMan0/GoPay/pkg/grpc_clients"
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
		syscall.SIGTERM,
		syscall.SIGINT,
	)
	defer cancel()

	slog.Info("Loading config...")
	cfg, err := config.Load()
	if err != nil {
		return fmt.Errorf("load config: %v", err)
	}
	slog.Info("Config loaded", slog.Any("config", cfg))

	slog.Info("Creating gRPC clients...")
	authGRPCClient, authConn, err := grpc_clients.NewAuthClient(cfg.AuthGRPCClient.Addr)
	if err != nil {
		return fmt.Errorf("create Auth gRPC client: %v", err)
	}
	defer func() {
		slog.Info("Closing Auth gRPC client connection...")
		closeErr := authConn.Close()
		if closeErr != nil {
			err = errors.Join(err, fmt.Errorf("close Auth gRPC client connection: %v", closeErr))
			return
		}

		slog.Info("Auth gRPC client connection closed")
	}()

	accountManagerGRPCClient, accountManagerConn, err := grpc_clients.NewAccountManagerClient(cfg.AccountManagerGRPCClient.Addr)
	if err != nil {
		return fmt.Errorf("create account manager gRPC client: %v", err)
	}
	defer func() {
		slog.Info("Closing Account Manager gRPC client connection...")
		closeErr := accountManagerConn.Close()
		if closeErr != nil {
			err = errors.Join(err, fmt.Errorf("close Account Manager gRPC client: %v", closeErr))
			return
		}

		slog.Info("Account Manager gRPC client connection closed")
	}()

	slog.Info("gRPC clients created")

	// Utils
	vldt := validator.New()

	accountManagerCL := grpc.NewAccountManagerClient(accountManagerGRPCClient)
	transferManagerCl := grpc.NewTransferManagerClient()
	operationFeedCl := grpc.NewOperationFeedClient()
	authCl := grpc.NewAuthServiceClient(authGRPCClient)

	// Controllers for different microservices
	accountManagerController := http.NewAccountManagerController(accountManagerCL)
	authServiceController := http.NewAuthServiceController(authCl, vldt)
	operationFeedController := http.NewOperationFeedController()
	transferManagerController := http.NewTransferManagerController()

	// General controller
	ctrl := controller.New(
		transferManagerController,
		operationFeedController,
		authServiceController,
		accountManagerController,
	)

	app := server.NewFiberApp(ctrl)
	errChan := make(chan error, 1)

	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()

		addr := fmt.Sprintf(":%s", cfg.HTTPServer.Addr)
		listenErr := app.Listen(addr)
		if listenErr != nil {
			slog.Error("listen error",
				slog.String("addr", cfg.HTTPServer.Addr),
				slog.String("error", listenErr.Error()),
			)

			errChan <- fmt.Errorf("listen error: %v", listenErr)
			cancel()
		}
	}()

	wg.Add(1)
	go func() {
		<-ctx.Done()

		slog.Info("Shutting down HTTP server...")
		shutDownErr := app.Shutdown()
		if shutDownErr != nil {
			slog.Error("HTTP server shutdown failed", slog.Any("error", shutDownErr))
			errChan <- fmt.Errorf("shutting down: %w", shutDownErr)
		} else {
			slog.Info("HTTP server stopped gracefully")
		}
	}()

	wg.Wait()
	close(errChan)
	for err = range errChan {
		if err != nil {
			err = errors.Join(err, fmt.Errorf("app shutdown: %w", err))
		}
	}
	return nil
}
