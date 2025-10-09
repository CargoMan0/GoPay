package grpc

import (
	"github.com/CargoMan0/GoPay/pkg/accountmanager"
)

type AccountManagerAdapter struct {
	cl accountmanager.AccountManagerClient
}

func NewAccountManagerAdapter(cl accountmanager.AccountManagerClient) *AccountManagerAdapter {
	return &AccountManagerAdapter{
		cl: cl,
	}
}
