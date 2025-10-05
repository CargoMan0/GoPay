package config

type Config struct {
	GRPCServer GRPCServer
}

type GRPCServer struct {
	Port int
}

func Load() (*Config, error) {
	var cfg Config

	return &cfg, nil
}
