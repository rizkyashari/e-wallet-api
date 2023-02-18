package handler

import s "e-wallet-api/internal/service"

type Handler struct {
	userService   s.UserService
	authService   s.AuthService
	walletService s.WalletService
	jwtService    s.JWTService
}

type HandlerConfig struct {
	UserService   s.UserService
	AuthService   s.AuthService
	WalletService s.WalletService
	JWTService    s.JWTService
}

func NewHandler(c *HandlerConfig) *Handler {
	return &Handler{
		userService:   c.UserService,
		authService:   c.AuthService,
		walletService: c.WalletService,
		jwtService:    c.JWTService,
	}
}
