package models

type NewAccountData struct {
	Username string
	Email    string
	Password string
}

type NewAccountResult struct {
	WalletAddress string
	AccessToken   string
	RefreshToken  string
}
