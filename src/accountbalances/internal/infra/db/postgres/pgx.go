package postgres

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"sync/atomic"
	"time"
)

type DBCluster struct {
	primary  *pgxpool.Pool
	replicas []*pgxpool.Pool
	// nextReplica field is user for Round-Robin
	nextReplica uint32
}

type DBConfig struct {
	Host       string
	User       string
	Password   string
	DB         string
	SSLMode    string
	Port       int
	IsReadOnly bool
}

func NewDBCluster(ctx context.Context, primaryCfg DBConfig, replicaCfgs []DBConfig) (*DBCluster, error) {
	primary, err := newPool(ctx, primaryCfg)
	if err != nil {
		return nil, fmt.Errorf("failed to create primary pool: %w", err)
	}

	var (
		replicas = make([]*pgxpool.Pool, 0, len(replicaCfgs))
		pool     *pgxpool.Pool
	)

	for _, cfg := range replicaCfgs {
		pool, err = newPool(ctx, cfg)
		if err != nil {
			return nil, fmt.Errorf("failed to create replica pool: %w", err)
		}
		replicas = append(replicas, pool)
	}

	return &DBCluster{
		primary:  primary,
		replicas: replicas,
	}, nil
}

func (d *DBCluster) GetPrimary() *pgxpool.Pool {
	return d.primary
}

func (d *DBCluster) GetReplica() *pgxpool.Pool {
	if len(d.replicas) == 0 {
		return d.primary
	}
	n := atomic.AddUint32(&d.nextReplica, 1)
	return d.replicas[(int(n)-1)%len(d.replicas)]
}

func newPool(ctx context.Context, cfg DBConfig) (*pgxpool.Pool, error) {
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DB, cfg.SSLMode)

	if cfg.IsReadOnly {
		connStr += " default_transaction_read_only=true"
	}

	config, err := pgxpool.ParseConfig(connStr)
	if err != nil {
		return nil, err
	}

	config.MaxConns = 35
	config.MinConns = 5
	config.MaxConnLifetime = 30 * time.Minute

	return pgxpool.NewWithConfig(ctx, config)
}
