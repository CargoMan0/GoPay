package config

type Config struct {
	HTTPServer               HTTPServer
	AuthGRPCClient           AuthGRPCClient
	AccountManagerGRPCClient AccountManagerGRPCClient
}

type HTTPServer struct {
	Addr string
}

type AuthGRPCClient struct {
	Addr string
}

type AccountManagerGRPCClient struct {
	Addr string
}

func Load() (*Config, error) {
	var cfg Config

	return &cfg, nil
}
