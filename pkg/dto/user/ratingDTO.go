package user

import (
	"dbPractice/pkg/constants"
	db2 "dbPractice/pkg/db"
	"log"
	"net/http"
)

func RatingDTO(w http.ResponseWriter, id string, rating int8) {
	db := db2.ConnectDB()
	_, dbErr := db.Exec(constants.IncreaseRatingByID, rating, id)
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
