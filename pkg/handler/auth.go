package handler

import (
	"fmt"
	"github.com/davitdarsalia/LendAppBackend/entities"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func (h *Handler) signUp(c *gin.Context) {
	var userInstance entities.User

	if err := c.BindJSON(&userInstance); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Authorization.RegisterUser(&userInstance)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	/* Writing statusCode to the response */
	c.JSON(http.StatusCreated, map[string]interface{}{
		"user_id":    id,
		"message":    "User Created Successfully",
		"created_at": time.Now().Format("2006-01-02 15:04:05"),
	})

	fmt.Println(userInstance)

}

func (h *Handler) signIn(c *gin.Context) {

}
