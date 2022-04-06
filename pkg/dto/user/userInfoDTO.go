package user

import (
	"dbPractice/pkg/constants"
	"dbPractice/pkg/db"
	"dbPractice/pkg/models"
	"log"
	"net/http"
)

func UserInfoDTO(w http.ResponseWriter, id string) models.Info {
	var info models.Info
	dB := db.ConnectDB()
	row := dB.QueryRow(constants.FetchUserInfo, id).Scan(&info.Balance, &info.Rating)

	if row != nil {
		errPayload := "0"
		w.WriteHeader(http.StatusNotFound)
		_, writeErr := w.Write([]byte(errPayload))

		if writeErr != nil {
			log.Fatal(writeErr)
		}
	}

	return info
}
