package user

import (
	"dbPractice/pkg/dto/user"
	"dbPractice/pkg/models"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func IncreaseRating(w http.ResponseWriter, r *http.Request) {
	var rating models.UserRating
	identifier := mux.Vars(r)
	id := identifier["id"]
	decodeErr := json.NewDecoder(r.Body).Decode(&rating)
	if decodeErr != nil {
		log.Println(decodeErr)
	}
	defer func() {
		closeErr := r.Body.Close()
		if closeErr != nil {
			log.Fatal(closeErr)
		}
	}()
	user.RatingDTO(w, id, rating.Rating)
}
