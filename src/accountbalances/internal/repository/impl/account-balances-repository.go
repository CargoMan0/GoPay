package impl

import (
	"context"
	"github.com/CargoMan0/GoPay/src/accountbalances/internal/core"
	"github.com/CargoMan0/GoPay/src/accountbalances/internal/infra/db/postgres"
	"github.com/google/uuid"
)

type accountBalanceRepositoryImpl struct {
	cluster *postgres.DBCluster
}

func NewAccountBalanceRepository(cluster *postgres.DBCluster) *accountBalanceRepositoryImpl {
	return &accountBalanceRepositoryImpl{
		cluster: cluster,
	}
}

func (a *accountBalanceRepositoryImpl) BatchInsertAccountBalance(ctx context.Context, balances []core.AccountBalance) error {
	//TODO implement me
	panic("implement me")
}

func (a *accountBalanceRepositoryImpl) InsertAccountBalanceAudit(ctx context.Context, audit core.AccountBalanceAudit) error {
	//TODO implement me
	panic("implement me")
}

func (a *accountBalanceRepositoryImpl) GetAllCurrencies(ctx context.Context) ([]core.Currency, error) {
	//TODO implement me
	panic("implement me")
}

func (a *accountBalanceRepositoryImpl) GetAccountBalances(ctx context.Context, userID uuid.UUID) ([]core.AccountBalance, error) {
	//TODO implement me
	panic("implement me")
}
