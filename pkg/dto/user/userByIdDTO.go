package user

import (
	"dbPractice/pkg/constants"
	"dbPractice/pkg/db"
	"dbPractice/pkg/models"
	"log"
	"net/http"
)

func UserByIdDTO(w http.ResponseWriter, id string) models.UserBasicInfo {
	var user models.UserBasicInfo
	dB := db.ConnectDB()
	row := dB.QueryRow(constants.GetUserByID, id).Scan(&user.Id, &user.Email, &user.FirstName, &user.LastName, &user.Age)
	if row != nil {
		w.WriteHeader(http.StatusNotFound)
		_, writeErr := w.Write([]byte("No User Found"))

		if writeErr != nil {
			log.Println(writeErr)
		}
	}
	return user
}
