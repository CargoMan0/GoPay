package grpc

import (
	"context"
	"github.com/CargoMan0/GoPay/gateway/internal/models"
	"github.com/google/uuid"
)

type TransferFeedAdapter struct {
}

func NewTransferFeedAdapter() *TransferFeedAdapter {
	return &TransferFeedAdapter{}
}

func (t *TransferFeedAdapter) NewTransfer(ctx context.Context, data *models.NewTransferData) (uuid.UUID, error) {

}

func (t *TransferFeedAdapter) CancelTransfer(ctx context.Context, id uuid.UUID) error {

}
