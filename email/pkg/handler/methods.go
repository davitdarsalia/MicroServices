package handler

import (
	"email/internal/entities"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) dummyMethod(c *gin.Context) {
	var e entities.Email

	if err := c.BindJSON(&e); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	h.service.Dummy()

	//if err != nil {
	//	newErrorResponse()
	//}

}
