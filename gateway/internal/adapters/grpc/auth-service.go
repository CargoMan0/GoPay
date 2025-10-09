package grpc

import (
	"context"
	"fmt"
	"github.com/CargoMan0/GoPay/gateway/internal/models"
	"github.com/CargoMan0/GoPay/pkg/gen/auth"
	"github.com/google/uuid"
)

type AuthServiceAdapter struct {
	cl auth.AuthServiceClient
}

func NewAuthServiceAdapter(client auth.AuthServiceClient) *AuthServiceAdapter {
	return &AuthServiceAdapter{
		cl: client,
	}
}

func (a *AuthServiceAdapter) Register(ctx context.Context, data *models.RegisterData) (*models.RegisterResult, error) {
	resp, err := a.cl.Register(ctx, &auth.RegisterRequest{
		Email:    data.Email,
		Password: data.Password,
		Username: data.Username,
	})
	if err != nil {
		return nil, fmt.Errorf("auth service gRPC client: register: %w", err)
	}

	userID, err := uuid.Parse(resp.UserId)
	if err != nil {
		return nil, fmt.Errorf("auth service gRPC client: parse user id: %w", err)
	}

	res := &models.RegisterResult{
		UserID:    userID,
		Email:     resp.Email,
		Username:  resp.Username,
		CreatedAt: resp.CreatedAt.AsTime(),
	}

	return res, nil
}

func (a *AuthServiceAdapter) Login(ctx context.Context, data *models.RegisterData) (*models.LoginResult, error) {
	resp, err := a.cl.Login(ctx, &auth.LoginRequest{
		Email:     data.Email,
		Password:  data.Password,
		IpAddress: data.IP,
		UserAgent: data.UserAgent,
	})
	if err != nil {
		return nil, fmt.Errorf("auth service gRPC client: login: %w", err)
	}

	res := &models.LoginResult{
		SessionID: resp.SessionId,
		ExpiresAt: resp.ExpiresAt.AsTime(),
	}

	return res, nil
}

func (a *AuthServiceAdapter) Logout(ctx context.Context, sessionID string) error {
	_, err := a.cl.Logout(ctx, &auth.LogoutRequest{
		SessionId: sessionID,
	})
	if err != nil {
		return fmt.Errorf("auth service gRPC client: logout: %w", err)
	}

	return nil
}

func (a *AuthServiceAdapter) ValidateSession(ctx context.Context, sessionID string) (*models.ValidateSessionResult, error) {
	resp, err := a.cl.ValidateSession(ctx, &auth.ValidateSessionRequest{
		SessionId: sessionID,
	})
	if err != nil {
		return nil, fmt.Errorf("auth service gRPC client: validate session: %w", err)
	}

	res := &models.ValidateSessionResult{
		Username:   resp.Username,
		Email:      resp.Email,
		ExpiresAt:  resp.ExpiresAt.AsTime(),
		LastUsedAt: resp.LastActivity.AsTime(),
	}

	return res, nil
}
