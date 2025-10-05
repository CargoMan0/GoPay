package config

type Config struct {
	GRPCServer   GRPCServer
	TokenManager TokenManager
}

type GRPCServer struct {
	Port int
}

type TokenManager struct {
	Secret string
}

func Load() (*Config, error) {
	var cfg Config

	return &cfg, nil
}
