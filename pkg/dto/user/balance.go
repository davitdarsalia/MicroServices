package user

import (
	"dbPractice/pkg/constants"
	db2 "dbPractice/pkg/db"
	"log"
	"net/http"
)

func BalanceDTO(w http.ResponseWriter, id string, balance int32) {
	db := db2.ConnectDB()

	_, dbErr := db.Exec(constants.IncreaseBalance, balance, id)

	if dbErr != nil {
		w.WriteHeader(http.StatusNotFound)
		log.Println(dbErr)
		return
	}

	w.WriteHeader(http.StatusResetContent)

	_, writeErr := w.Write([]byte("Balance Updated"))

	if writeErr != nil {
		log.Println(writeErr)
	}
}
