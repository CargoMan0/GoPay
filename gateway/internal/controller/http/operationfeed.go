package http

import "github.com/gin-gonic/gin"

type OperationFeedController struct {
}

func NewOperationFeedController() *OperationFeedController {
	return &OperationFeedController{}
}

func (c *OperationFeedController) GetTransfer(ctx *gin.Context) {}

func (c *OperationFeedController) GetTransfers(ctx *gin.Context) {}
