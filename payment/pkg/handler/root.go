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
		testPayments.POST("/fetch-public-key", h.FetchPublicKey)
	}

	realPayments := r.Group("api/payment")
	{
		realPayments.POST("/fetch-public-key")
	}

	return r
}

// New - Returns New And Only Instance Of Handler
func New(s *service.Service) *Handler {
	return &Handler{services: s}
}
