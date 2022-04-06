package models

type Info struct {
	Balance int32 `json:"balance"`
	Rating  int8  `json:"rating"`
}

type UserBasicInfo struct {
	Id        string `json:"id"`
	Email     string `json:"email"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Age       int8   `json:"age"`
}

type UserRating struct {
	Rating int8 `json:"rating"`
}

type UserBalance struct {
	Balance int32 `json:"balance"`
}
