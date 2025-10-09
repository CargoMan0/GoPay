package cache

import "github.com/redis/go-redis/v9"

type Config struct {
	Addr string
}

func NewRedis(c Config) *redis.Client {
	cl := redis.NewClient(&redis.Options{
		Addr: c.Addr,
	})

	return cl
}
