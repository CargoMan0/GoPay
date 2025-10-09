package grpc

import (
	"context"
	"github.com/google/uuid"
)

type TransferManagerAdapter struct {
}

func NewTransferManagerAdapter() *TransferManagerAdapter {
	return &TransferManagerAdapter{}
}

func (t *TransferManagerAdapter) NewTransfer(ctx context.Context) (uuid.UUID, error) {
	return uuid.Nil, nil
}

func (t *TransferManagerAdapter) CancelTransfer(ctx context.Context, id uuid.UUID) error {
	return nil
}
