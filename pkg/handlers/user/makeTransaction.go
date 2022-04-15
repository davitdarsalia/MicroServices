package user

import (
	"dbPractice/pkg/dto/user"
	"dbPractice/pkg/models"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

func MakeTransaction(w http.ResponseWriter, r *http.Request) {
	var transactionModel models.UserTransaction
	identifier := mux.Vars(r)
	id := identifier["id"]
	decodeErr := json.NewDecoder(r.Body).Decode(&transactionModel)
	if decodeErr != nil {
		log.Println(decodeErr)
	}
	defer func() {
		closeErr := r.Body.Close()
		if closeErr != nil {
			log.Fatal(closeErr)
		}
	}()
	idInt, parseError := strconv.ParseInt(id, 0, 0)
	if parseError != nil {
		log.Fatal(parseError, "DDD")
	}
	user.TransactionDTO(w, transactionModel, idInt)
}
