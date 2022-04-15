package global

import (
	"dbPractice/pkg/dto/global"
	"encoding/json"
	"log"
	"net/http"
)

func GetAllCountryCodes(w http.ResponseWriter, r *http.Request) {
	codes, err := global.GetCountryCodesDTO(w)
	if err != nil {
		log.Fatal(err)
	}
	codesJSON, err := json.Marshal(codes)
	_, writeError := w.Write(codesJSON)
	if writeError != nil {
		log.Fatal(writeError)
	}
}

func GetCountryCodeByID(w http.ResponseWriter, r *http.Request) {

}
