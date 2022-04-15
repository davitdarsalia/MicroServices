package global

import (
	"dbPractice/pkg/constants"
	db2 "dbPractice/pkg/db"
	"dbPractice/pkg/models"
	"log"
	"net/http"
)

func GetCountryCodesDTO(w http.ResponseWriter) ([]models.CountryModel, error) {
	var allCodes []models.CountryModel
	db := db2.ConnectDB()
	rows, err := db.Query(constants.GetAllCountryCodes)
	defer func() {
		rowCloseErr := rows.Close()
		if rowCloseErr != nil {
			log.Fatal(rowCloseErr)
		}
	}()
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		var country models.CountryModel
		if scanError := rows.Scan(); scanError != nil {
			log.Println(scanError)
		}
		allCodes = append(allCodes, country)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	return allCodes, nil
}

func GetCountryByIdDTO() {
	return
}
