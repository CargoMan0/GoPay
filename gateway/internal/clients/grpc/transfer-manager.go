package grpc

import (
	"context"
	"github.com/google/uuid"
)

type TransferFeedClient struct {
}

func NewTransferFeedAdapter() *TransferFeedClient {
	return &TransferFeedClient{}
}

func (t *TransferFeedClient) NewTransfer(ctx context.Context, data *models.NewTransferData) (uuid.UUID, error) {

}

func (t *TransferFeedClient) CancelTransfer(ctx context.Context, id uuid.UUID) error {

}
