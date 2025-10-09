package http

type registerRequest struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type registerResponse struct {
	Email     string `json:"email"`
	Username  string `json:"username"`
	UserID    string `json:"session_id"`
	CreatedAt string `json:"created_at"`
}

type loginResponse struct {
	SessionID string `json:"session_id"`
	ExpiresAt string `json:"expires_at"`
}

type logoutRequest struct {
	SessionID string `json:"session_id"`
}
