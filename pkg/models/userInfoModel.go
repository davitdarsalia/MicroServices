package models

type UserInfoModel struct {
	balance int32
	rating  int8
}

type UserBasicInfo struct {
	Id        string `json:"id"`
	Email     string `json:"email"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Age       int8   `json:"age"`
}
