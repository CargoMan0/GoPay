package http

import (
	"github.com/CargoMan0/GoPay/gateway/internal/adapter/grpc"
	"github.com/gin-gonic/gin"
	"github.com/shopspring/decimal"
)

type TransferManagerController struct {
	grpcAdapter grpc.TransferFeedAdapter
}

func NewTransferManagerController() *TransferManagerController {
	return &TransferManagerController{}

}

func (t *TransferManagerController) NewTransfer(ctx *gin.Context) {
	var req newTransferRequestDTO

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "invalid request", "details": err.Error()})
		return
	}

	decimal.NewFromFloat(req.Amount)

	id, err := t.grpcAdapter.NewTransfer(ctx)
	if err != nil {
		handleError(ctx, err)
	}

	ctx.JSON(200, gin.H{
		"id": id,
	})
}

func (t *TransferManagerController) CancelTransfer(ctx *gin.Context) {}
