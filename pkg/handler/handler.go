package handler

import (
	"github.com/davitdarsalia/LendAppBackend/pkg/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func (h *Handler) InitRoutes() *gin.Engine {
	r := gin.New()

	auth := r.Group("/api/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}

	account := r.Group("/api/account")
	{
		account.GET("/user-info", h.signUp)
		account.GET("/rating", h.signUp)
		account.GET("/bonus", h.signUp)
		account.GET("/balance", h.signUp)
		account.GET("/cashback", h.signUp)
		account.GET("/deposit", h.signUp)
	}

	transactions := r.Group("/api/transactions")
	{
		transactions.PUT("/rating", h.signUp)
		transactions.PUT("/deposit", h.signUp)
		transactions.PUT("/balance", h.signUp)
		transactions.PUT("/available-currencies", h.signUp)
		transactions.PUT("/cashback", h.signUp)
		transactions.PUT("/bonus", h.signUp)
	}

	deletions := r.Group("/api/remove")
	{
		deletions.DELETE("/deposit", h.signUp)
		deletions.DELETE("/currency", h.signUp)
	}

	return r
}
