package main

import (
	"context"
	"fmt"
	"github.com/CargoMan0/GoPay/gateway/internal/config"
	"log"
	"log/slog"
	"os/signal"
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

	slog.Info("Creating gRPC client...")

	slog.Info("gRPC client created")

	<-ctx.Done()
	return nil
}
