package utils

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/sirupsen/logrus"
)

type localError struct {
	Message string `json:"message"`
}

func Error(c *gin.Context, statusCode int, message string) {
	logrus.Errorf(message)
	c.AbortWithStatusJSON(statusCode, localError{Message: message})
}

func PgxErrorHandler(err error) {
	if err != nil {
		var pgErr *pgconn.PgError

		if errors.As(err, &pgErr) {
			fmt.Println(pgErr.Message) // => syntax error at end of input
			fmt.Println(pgErr.Code)    // => 42601
		}
	}
}

// CreateBenchmarkDocs - TODO
func CreateBenchmarkDocs() {

}
