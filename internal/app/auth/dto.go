package auth

type RegisterRequest struct {
	Username string `json:"username" binding:"required,min=3"`
	Password string `json:"password" binding:"required,min=6"`
}

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type CheckUsernameRequest struct {
	Username string `json:"username" binding:"required,min=3"`
}

type AuthResponse struct {
	Success bool     `json:"success"`
	Data    AuthData `json:"data"`
}

type CheckUsernameResponse struct {
	Success bool                   `json:"success"`
	Data    CheckUsernameData     `json:"data"`
}

type AuthData struct {
	UserID int    `json:"userId"`
	Token  string `json:"token"`
}

type CheckUsernameData struct {
	Available bool   `json:"available"`
	Message   string `json:"message"`
}