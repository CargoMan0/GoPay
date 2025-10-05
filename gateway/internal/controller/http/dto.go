package http

type newTransferRequestDTO struct {
	FromAccount string  `json:"from_account" binding:"required"`
	ToAccount   string  `json:"to_account" binding:"required"`
	Amount      float64 `json:"amount" binding:"required,gt=0"`
}
