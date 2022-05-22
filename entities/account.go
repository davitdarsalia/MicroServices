package entities

type ProfileDetails struct {
	ProfileImage       []byte `json:"profileimage"`
	Followers          int    `json:"followers"`
	Following          int    `json:"following"`
	BlockedUsersAmount int    `json:"blocked_users_amount"`
	WorkingPlace       string `json:"working_place"`
	Education          string `json:"education"`
	Origin             string `json:"origin"`
	AdditionalEmail    string `json:"additional_email"`
	UserID             string `json:"userid"`
}

type UserInfo struct {
	User
	ProfileImage       []byte `json:"profileimage"`
	Followers          int    `json:"followers"`
	Following          int    `json:"following"`
	BlockedUsersAmount int    `json:"blocked_users_amount"`
	WorkingPlace       string `json:"working_place"`
	Education          string `json:"education"`
	Origin             string `json:"origin"`
	AdditionalEmail    string `json:"additional_email"`
	UserID             string `json:"userid"`
}
