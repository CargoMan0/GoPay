package config

type Config struct {
	GRPCServer   GRPCServer
	TokenManager TokenManager
	EventSender  EvenSender
	Producer     Producer
}

type Producer struct {
}

type GRPCServer struct {
	Port int
}

type TokenManager struct {
	Secret string
}

type EvenSender struct {
	HandlePeriodSeconds int
	MaxBatchSize        uint8
}

func Load() (*Config, error) {
	var cfg Config

	return &cfg, nil
}
