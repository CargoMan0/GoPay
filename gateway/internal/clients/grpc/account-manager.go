package grpc

import (
	"context"
	"fmt"
	"github/com/CargoMan0/GoPay/gateway/internal/models"
	"github/com/CargoMan0/GoPay/pkg/accountmanager"
)

type AccountManagerClient struct {
	cl accountmanager.AccountManagerClient
}

func NewAccountManagerClient(cl accountmanager.AccountManagerClient) *AccountManagerClient {
	return &AccountManagerClient{
		cl: cl,
	}
}

func (a *AccountManagerClient) NewAccount(ctx context.Context, data *models.NewAccountData) (*models.NewAccountResult, error) {
	req := &accountmanager.NewAccountRequest{
		Username: data.Username,
		Password: data.Password,
		Email:    data.Email,
	}

	account, err := a.cl.NewAccount(ctx, req)
	if err != nil {
		// TODO: handle errors from grpc
		return nil, fmt.Errorf("new account: %w", err)
	}

	resp := &models.NewAccountResult{
		WalletAddress: account.AccountAddress,
		AccessToken:   account.AccessToken,
		RefreshToken:  account.RefreshToken,
	}

	return resp, nil
}
