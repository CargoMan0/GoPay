package main

import (
	"context"
	"fmt"
	"github.com/CargoMan0/GoPay/src/accountbalances/internal/config"
	"github.com/CargoMan0/GoPay/src/accountbalances/internal/infra/broker/kafka"
	"github.com/CargoMan0/GoPay/src/accountbalances/internal/infra/db/postgres"
	"github.com/CargoMan0/GoPay/src/accountbalances/internal/outbox"
	"github.com/CargoMan0/GoPay/src/accountbalances/internal/repository/impl"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	err := run()
	if err != nil {
		log.Fatalf("run() returned error: %v", err)
	}
}

func run() error {
	ctx, cancel := signal.NotifyContext(
		context.Background(),
		os.Interrupt,
		os.Kill,
		syscall.SIGTERM,
	)
	defer cancel()

	cfg, err := config.Load()
	if err != nil {
		return fmt.Errorf("load config: %w", err)
	}

	dbCluster, err := newDBCluster(ctx, cfg.DatabaseCluster)
	if err != nil {
		return err
	}

	accountBalancesRepository := impl.NewAccountBalanceRepository(dbCluster)

	outboxEventSender := outbox.NewEventSender()
	go func() {
		outboxEventSender.Run(ctx)
	}()

	return nil
}

func newDBCluster(ctx context.Context, cfg config.DatabaseCluster) (*postgres.DBCluster, error) {
	cluster, err := postgres.NewDBCluster(ctx,
		postgres.DBConfig{
			Host:       cfg.Master.Host,
			Port:       cfg.Master.Port,
			User:       cfg.Master.User,
			DB:         cfg.Master.Name,
			Password:   cfg.Master.Password,
			SSLMode:    cfg.Master.SSLMode,
			IsReadOnly: false,
		},
		[]postgres.DBConfig{
			{
				Host:       cfg.Slave1.Host,
				Port:       cfg.Slave1.Port,
				User:       cfg.Slave1.User,
				DB:         cfg.Slave1.Name,
				Password:   cfg.Slave1.Password,
				SSLMode:    cfg.Slave1.SSLMode,
				IsReadOnly: true,
			}, {
				Host:       cfg.Slave2.Host,
				Port:       cfg.Slave2.Port,
				User:       cfg.Slave2.User,
				DB:         cfg.Slave2.Name,
				Password:   cfg.Slave2.Password,
				SSLMode:    cfg.Slave2.SSLMode,
				IsReadOnly: true,
			},
		},
	)

	return cluster, nil
}

func newEventPublisher(cfg config.ProducerConfig) {
	kafka.NewProducer([]string{})
}

func newEventReader(cfg config.ReaderConfig) {
	kafka.NewConsumer()
}
