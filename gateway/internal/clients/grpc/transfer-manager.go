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

func (t *TransferFeedClient) NewTransfer(ctx context.Context) (uuid.UUID, error) {
	return uuid.Nil, nil
}

func (t *TransferFeedClient) CancelTransfer(ctx context.Context, id uuid.UUID) error {
	return nil
}
