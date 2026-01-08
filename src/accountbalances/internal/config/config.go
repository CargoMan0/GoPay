package config

import (
	"fmt"
	"github.com/caarlos0/env/v11"
)

// Config struct represents application configuration.
type Config struct {
	// Database stores configuration for a database.
	DatabaseCluster DatabaseCluster `envPrefix:"DATABASE_CLUSTER_"`
	// Cache stores configuration for a cache.
	Cache CacheConfig `envPrefix:"CACHE_"`
}

// Load loads application config and returns a pointer to Config.
func Load() (*Config, error) {
	cfg, err := env.ParseAs[Config]()
	if err != nil {
		return nil, fmt.Errorf("failed to parse config as %T: %w", Config{}, err)
	}

	return &cfg, nil
}

type DatabaseCluster struct {
	// Master is for write operations
	Master DatabaseConfig `envPrefix:"MASTER_"`
	// Slaves is read-only database instances
	Slave1 DatabaseConfig `envPrefix:"SLAVE_1_"`
	Slave2 DatabaseConfig `envPrefix:"SLAVE_2_"`
}

type DatabaseConfig struct {
	Host     string `env:"HOST"`
	User     string `env:"USER"`
	Password string `env:"PASSWORD"`
	Name     string `env:"NAME"`
	SSLMode  string `env:"SSL_MODE"`
	Port     int    `env:"PORT"`
}

type CacheConfig struct {
	Host string `env:"HOST"`
	Port int    `env:"PORT"`
}

type ProducerConfig struct{}

type ReaderConfig struct{}
