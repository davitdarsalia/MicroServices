package user

import (
	"dbPractice/pkg/handlers/auth"
	"net/http"
)

func TransactionsHandler(w http.ResponseWriter, r *http.Request) {
	if auth.IsAuthorized == false {
		panic(any("Not Authorized"))
	}

}
