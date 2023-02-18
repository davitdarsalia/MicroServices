package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

const (
	authHeader = "Authorization"
	userCTX    = "userId"
)

func (h *Handler) authCheck(c *gin.Context) {
	header := c.GetHeader(authHeader)

	if header == "" {
		newErrorResponse(c, http.StatusUnauthorized, "Empty Auth Header")
		return
	}

	splitHeader := strings.Split(header, " ")

	if len(splitHeader) != 2 {
		newErrorResponse(c, http.StatusUnauthorized, "Invalid Auth Header")
		return
	}

	userID, err := h.service.Authorizer.CheckToken(splitHeader[1], "e021kr01k2t04k")

	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	c.Set(userCTX, userID)
}
