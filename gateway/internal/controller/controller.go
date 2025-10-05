package controller

import (
	"github.com/CargoMan0/GoPay/gateway/internal/controller/http"
	"github.com/gin-gonic/gin"
)

type Controller interface {
	SetupRoutes(e *gin.Engine)
}

type controller struct {
	transferManagerCtrl *http.TransferManagerController
	operationFeedCtrl   *http.OperationFeedController
	accountManagerCtrl *http.
}

func New(
	tmc *http.TransferManagerController,
	ofc *http.OperationFeedController,
) Controller {
	return &controller{
		transferManagerCtrl: tmc,
		operationFeedCtrl:   ofc,
	}
}

func (c *controller) SetupRoutes(eng *gin.Engine) {
	eng.Use(gin.Recovery())

	apiV1 := eng.Group("/v1")

	transferGroup := apiV1.Group("/transfer")
	transferGroup.GET("/:id", c.operationFeedCtrl.GetTransfer)
	transferGroup.GET("/", c.operationFeedCtrl.GetTransfers)
	transferGroup.POST("/", c.transferManagerCtrl.NewTransfer)
	transferGroup.DELETE("/:id", c.transferManagerCtrl.CancelTransfer)

	accountGroup := apiV1.Group("/account")
	accountGroup.GET("/:id")
	accountGroup.POST("/")

}
