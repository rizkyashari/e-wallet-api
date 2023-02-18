package main

import (
	"fmt"
	"os"

	"e-wallet-api/config"
	"e-wallet-api/internal/handler"
	"e-wallet-api/internal/repository"
	"e-wallet-api/internal/route"
	"e-wallet-api/internal/service"

	"github.com/gin-gonic/gin"
)

func main() {
	db := config.GetConn()

	userRepository := repository.NewUserRepository(&repository.URConfig{DB: db})
	walletRepository := repository.NewWalletRepository(&repository.WRConfig{DB: db})
	passwordResetRepository := repository.NewPasswordResetRepository(&repository.PRConfig{DB: db})
	sourceOfFundRepository := repository.NewSourceOfFundRepository(&repository.SRConfig{DB: db})
	transactionRepository := repository.NewTransactionRepository(&repository.TRConfig{DB: db})

	userService := service.NewUserService(&service.USConfig{UserRepository: userRepository})
	authService := service.NewAuthService(&service.ASConfig{UserRepository: userRepository, PasswordResetRepository: passwordResetRepository})
	walletService := service.NewWalletService(&service.WSConfig{UserRepository: userRepository, WalletRepository: walletRepository})
	transactionService := service.NewTransactionService(&service.TSConfig{TransactionRepository: transactionRepository, SourceOfFundRepository: sourceOfFundRepository, WalletRepository: walletRepository})
	jwtService := service.NewJWTService(&service.JWTSConfig{})

	h := handler.NewHandler(&handler.HandlerConfig{
		UserService:        userService,
		AuthService:        authService,
		WalletService:      walletService,
		TransactionService: transactionService,
		JWTService:         jwtService,
	})

	routes := route.NewRouter(&route.RouterConfig{UserService: userService, JWTService: jwtService})

	router := gin.Default()

	version := os.Getenv("API_VERSION")
	api := router.Group(fmt.Sprintf("/api/%s", version))

	routes.Auth(api, h)
	routes.Transaction(api, h)

	router.Run(":8000")
}
