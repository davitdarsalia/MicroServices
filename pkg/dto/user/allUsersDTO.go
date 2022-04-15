package user

import (
	"dbPractice/pkg/constants"
	"dbPractice/pkg/db"
	"dbPractice/pkg/models"
	"log"
	"net/http"
)

func AllUsersDTO(w http.ResponseWriter) ([]models.UserBasicInfo, error) {
	var allUsers []models.UserBasicInfo
	dB := db.ConnectDB()
	rows, err := dB.Query(constants.GetAllUsers)
	defer func() {
		rowCloseErr := rows.Close()
		if rowCloseErr != nil {
			log.Fatal(rowCloseErr)
		}
	}()
	if err != nil {
		log.Println(err)
		return nil, nil
	}
	for rows.Next() {
		var user models.UserBasicInfo

		if scanError := rows.Scan(&user.Id, &user.Email, &user.FirstName, &user.LastName, &user.Age); scanError != nil {
			log.Println(scanError)
		}
		allUsers = append(allUsers, user)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	return allUsers, nil
}
