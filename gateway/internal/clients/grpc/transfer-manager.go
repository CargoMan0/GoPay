package grpc

import (
	"context"
	"github.com/google/uuid"
)

type TransferManagerClient struct {
}

func NewTransferManagerClient() *TransferManagerClient {
	return &TransferManagerClient{}
}

func (t *TransferManagerClient) NewTransfer(ctx context.Context) (uuid.UUID, error) {
	return uuid.Nil, nil
}

func (t *TransferManagerClient) CancelTransfer(ctx context.Context, id uuid.UUID) error {
	return nil
}
