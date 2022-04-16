package user

import (
	"dbPractice/pkg/dto/user"
	"dbPractice/pkg/models"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func IncreaseBalance(w http.ResponseWriter, r *http.Request) {
	var balance models.UserBalance
	identifier := mux.Vars(r)
	id := identifier["id"]
	decodeErr := json.NewDecoder(r.Body).Decode(&balance)
	if decodeErr != nil {
		log.Println(decodeErr)
	}
	user.BalanceDTO(w, id, balance.Balance)
}
