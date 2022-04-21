package routes

import (
	"github.com/gin-gonic/gin"
	"log"
	"os"
)

func RootRouter() {
	defPort := os.Getenv("DEFAULT_PORT")
	r := gin.Default()
	/* AuthRoutes */
	AuthRoutes(r)
	/* Protected Routes */
	serverErr := r.Run(defPort)
	if serverErr != nil {
		log.Fatal(serverErr)
	}
}
