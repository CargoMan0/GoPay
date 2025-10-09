package grpc_clients

import (
	"fmt"
	"github/com/CargoMan0/GoPay/pkg/accountmanager"
	"github/com/CargoMan0/GoPay/pkg/gen/auth"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

/*
func NewOperationFeedClient(addr string) (currency.CurrencyServiceClient, *grpc.ClientConn, error) {
	dialOptions := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}

	conn, err := grpc.NewClient(addr, dialOptions...)
	if err != nil {
		return nil, nil, fmt.Errorf("grpc.NewClient: %w", err)
	}

	client := currency.NewCurrencyServiceClient(conn)

	return client, conn, nil
}
*/

func NewAuthClient(addr string) (auth.AuthServiceClient, *grpc.ClientConn, error) {
	dialOptions := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}

	conn, err := grpc.NewClient(addr, dialOptions...)
	if err != nil {
		return nil, nil, fmt.Errorf("grpc.NewClient: %w", err)
	}

	client := auth.NewAuthServiceClient(conn)

	return client, conn, nil
}

func NewAccountManagerClient(addr string) (accountmanager.AccountManagerClient, *grpc.ClientConn, error) {
	dialOptions := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}

	conn, err := grpc.NewClient(addr, dialOptions...)
	if err != nil {
		return nil, nil, fmt.Errorf("grpc.NewClient: %w", err)
	}

	client := accountmanager.NewAccountManagerClient(conn)

	return client, conn, nil
}
