package handler

import (
	"email/pkg/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *service.Service
}

func New(s *service.Service) *Handler {
	return &Handler{service: s}
}

func (h *Handler) DefineRoutes() *gin.Engine {
	r := gin.Default()

	email := r.Group("/dummy")
	{
		email.GET("dummy", h.dummyMethod)
	}

	return r
}
