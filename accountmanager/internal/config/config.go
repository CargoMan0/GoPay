package config

import (
	"flag"
	"fmt"
	"github.com/spf13/viper"
)

type Config struct {
	GRPCServer   GRPCServer   `mapstructure:"grpc_server"`
	TokenManager TokenManager `mapstructure:"token_manager"`
	EventSender  EvenSender   `mapstructure:"event_sender"`
	Producer     Producer     `mapstructure:"producer"`
}

type Producer struct {
}

type GRPCServer struct {
	Port int `mapstructure:"port"`
}

type TokenManager struct {
	Secret string `mapstructure:"secret"`
}

type EvenSender struct {
	HandlePeriodSeconds uint8 `mapstructure:"handle_period_seconds"`
	MaxBatchSize        uint8 `mapstructure:"max_batch_size"`
}

func Load() (Config, error) {
	path := getPath()

	var cfg Config
	viper.SetConfigFile(path)

	if err := viper.ReadInConfig(); err != nil {
		return cfg, fmt.Errorf("error reading config file: %w", err)
	}

	if err := viper.Unmarshal(&cfg); err != nil {
		return cfg, fmt.Errorf("unable to unmarshal config: %w", err)
	}

	return cfg, nil
}

func getPath() string {
	configPath := flag.String("config", "./config", "path to the config file")
	flag.Parse()

	return *configPath
}
