package config

type Config struct {
}

func Load() (*Config, error) {
	var cfg Config

	return &cfg, nil
}
