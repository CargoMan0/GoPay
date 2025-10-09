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

type registerRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}
