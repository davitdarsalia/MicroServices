package user

import (
	"dbPractice/pkg/dto/user"
	"dbPractice/pkg/models"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func IncreaseRating(w http.ResponseWriter, r *http.Request) {
	var rating models.UserRating

	identifier := mux.Vars(r)
	id := identifier["id"]

	decodeErr := json.NewDecoder(r.Body).Decode(&rating)

	fmt.Println(rating, "rating")

	if decodeErr != nil {
		log.Println(decodeErr)
	}

	user.RatingDTO(w, id, rating.Rating)

}
