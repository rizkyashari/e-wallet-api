package route

import (
	"e-wallet-api/internal/handler"

	"github.com/gin-gonic/gin"
)

func (r *Router) Auth(route *gin.RouterGroup, h *handler.Handler) {
	route.POST("/register", h.Register)
	route.POST("/login", h.Login)
	route.POST("/forgot-password", h.ForgotPassword)
	route.POST("/reset-password", h.ResetPassword)
}
