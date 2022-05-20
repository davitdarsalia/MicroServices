package entities

type ProfileDetails struct {
	ProfileImage       []byte `json:"profileimage" binding:"required"`
	Followers          int    `json:"followers" binding:"required"`
	Following          int    `json:"following" binding:"required"`
	BlockedUsersAmount int    `json:"blocked_users_amount" binding:"required"`
	WorkingPlace       string `json:"working_place" binding:"required"`
	Education          string `json:"education" binding:"required"`
	Origin             string `json:"origin" binding:"required"`
	AdditionalEmail    string `json:"additional_email" binding:"required"`
	UserID             string `json:"userid" db:"userid"`
}
