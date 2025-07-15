package auth

type RegisterRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
	Nickname string `json:"nickname" binding:"required,min=2"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type AuthResponse struct {
	Success bool     `json:"success"`
	Data    AuthData `json:"data"`
}

type AuthData struct {
	UserID   int    `json:"userId"`
	Nickname string `json:"nickname,omitempty"`
	Token    string `json:"token"`
}