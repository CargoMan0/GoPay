package grpc

import (
	"context"
	"fmt"
	"github/com/CargoMan0/GoPay/gateway/internal/models"
	"github/com/CargoMan0/GoPay/pkg/gen/auth"
)

type AuthServiceClient struct {
	cl auth.AuthServiceClient
}

func NewAuthServiceClient(client auth.AuthServiceClient) *AuthServiceClient {
	return &AuthServiceClient{
		cl: client,
	}
}

func (a *AuthServiceClient) Register(ctx context.Context, data *models.RegisterData) (string, error) {
	resp, err := a.cl.Register(ctx, &auth.RegisterRequest{
		Email:    data.Email,
		Password: data.Password,
		Username: data.Username,
	})
	if err != nil {
		return "", fmt.Errorf("auth service gRPC client: register: %w", err)
	}

	_ = resp
	return "", nil
}

func (a *AuthServiceClient) Login(ctx context.Context, data *models.RegisterData) (string, error) {
	resp, err := a.cl.Login(ctx, &auth.LoginRequest{
		Email:     data.Email,
		Password:  data.Password,
		IpAddress: data.IP,
		UserAgent: data.UserAgent,
	})
	if err != nil {
		return "", fmt.Errorf("auth service gRPC client: login: %w", err)
	}

	_ = resp

	return resp.SessionId, nil
}

func (a *AuthServiceClient) Logout(ctx context.Context, sessionID string) error {
	_, err := a.cl.Logout(ctx, &auth.LogoutRequest{
		SessionId: sessionID,
	})
	if err != nil {
		return fmt.Errorf("auth service gRPC client: logout: %w", err)
	}

	return nil
}

func (a *AuthServiceClient) ValidateSession(ctx context.Context, sessionID string) (string, error) {
	resp, err := a.cl.ValidateSession(ctx, &auth.ValidateSessionRequest{
		SessionId: sessionID,
	})
	if err != nil {
		return "", fmt.Errorf("auth service gRPC client: validate session: %w", err)
	}

	_ = resp

	return "", nil
}
