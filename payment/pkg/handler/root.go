package handler

import (
	"github.com/davitdarsalia/payment/pkg/service"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Handler struct {
	services *service.Service
}

func (h *Handler) Routes() *gin.Engine {
	r := gin.Default()

	docs := r.Group("/docs")
	{
		docs.GET("/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	testPayments := r.Group("api/payment/test")
	{
		testPayments.POST("/create-user", h.FetchPublicKey)
		testPayments.POST("/login")
		testPayments.POST("/refresh-login")

		testPayments.POST("/reset-password")

		testPayments.POST("/verify-reset-email")
	}

	realPayments := r.Group("api/payment")
	{
		realPayments.POST("/create-user")
		realPayments.POST("/login")
		realPayments.POST("/refresh-login")

		realPayments.POST("/reset-password")
	}

	return r
}

// New - Returns New And Only Instance Of Handler
func New(s *service.Service) *Handler {
	return &Handler{services: s}
}
