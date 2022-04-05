package user

import (
	"fmt"
	"net/http"
	"strings"
)

func TransactionsHandler(w http.ResponseWriter, r *http.Request) {

	token := r.Header.Get("token")

	fmt.Println("Done")

	if strings.HasPrefix(token, "Bearer") {
		// Trim space
		token = strings.TrimPrefix(token, "Bearer")

	}

}
