package config

type Config struct {
	HTTPServer               HTTPServer               `mapstructure:"http_server"`
	Redis                    Redis                    `mapstructure:"redis"`
	AuthGRPCClient           AuthGRPCClient           `mapstructure:"auth_grpc_client"`
	AccountManagerGRPCClient AccountManagerGRPCClient `mapstructure:"account_manager_grpc_client"`
	AccountBalanceGRPCClient AccountBalanceGRPCClient `mapstructure:"account_balance_grpc_client"`
}

type HTTPServer struct {
	Addr         string `mapstructure:"addr"`
	ReadTimeout  uint   `mapstructure:"read_timeout"`
	WriteTimeout uint   `mapstructure:"write_timeout"`
}

type AuthGRPCClient struct {
	Addr string `mapstructure:"addr"`
}

type AccountManagerGRPCClient struct {
	Addr string `mapstructure:"addr"`
}

type AccountBalanceGRPCClient struct {
	Addr string `mapstructure:"addr"`
}

type Redis struct {
	Addr string `mapstructure:"addr"`
}

func Load() (*Config, error) {
	var cfg Config

	return &cfg, nil
}
