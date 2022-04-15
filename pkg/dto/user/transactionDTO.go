package user

import (
	"dbPractice/pkg/constants"
	db2 "dbPractice/pkg/db"
	"dbPractice/pkg/models"
	"log"
	"net/http"
)

func TransactionDTO(w http.ResponseWriter, t models.UserTransaction, id int64) {
	db := db2.ConnectDB()
	_, dbErr := db.Exec(constants.RegisterTransaction, id, t.Recipient, t.Amount, t.Currency)
	if dbErr != nil {
		w.WriteHeader(http.StatusNotFound)
		log.Println(dbErr)
		return
	}
	w.WriteHeader(http.StatusResetContent)
	_, writeErr := w.Write([]byte("Rating Updated"))
	if writeErr != nil {
		log.Println(writeErr)
	}
}
