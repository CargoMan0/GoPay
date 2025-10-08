package http

type newAccountRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type newAccountResponse struct {
	WalletAddress string `json:"wallet_address"`
	AccessToken   string `json:"access_token"`
	RefreshToken  string `json:"refresh_token"`
}

type newTransferRequestDTO struct {
	FromAccount string  `json:"from_account" binding:"required"`
	ToAccount   string  `json:"to_account" binding:"required"`
	Amount      float64 `json:"amount" binding:"required,gt=0"`
}
