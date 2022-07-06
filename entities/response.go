package entities

type RegisteredUserResponse struct {
	UserId    int    `json:"user_id"`
	Message   string `json:"message"`
	CreatedAt string `json:"created_at"`
}

type SignedInUserResponse struct {
	Message         string `json:"message"`
	AccessToken     string `json:"access_token"`
	AccessTokenExp  string `json:"access_token_exp"`
	RefreshToken    string `json:"refresh_token"`
	RefreshTokenExp string `json:"refresh_token_exp"`
}

type ResetPasswordResponse struct {
	UserID  string `json:"user_id"`
	Message string `json:"message"`
}

type ValidateResetPasswordResponse struct {
	Message   string `json:"message"`
	ResetDate string `json:"reset_date"`
}

type ResetPasswordProfileResponse struct {
	Message   string `json:"message"`
	ResetDate string `json:"reset_date"`
}

type GetProfileDetailsResponse struct {
	Message            string `json:"message"`
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
type GetUserInfoResponse struct {
	User
	Message            string `json:"message"`
	ProfileImage       []byte `json:"profileimage"`
	Followers          int    `json:"followers"`
	Following          int    `json:"following"`
	BlockedUsersAmount int    `json:"blocked_users_amount"`
	WorkingPlace       string `json:"working_place"`
	Education          string `json:"education"`
	Origin             string `json:"origin"`
	AdditionalEmail    string `json:"additional_email"`
}

type GetTrustedDevices struct {
	Message    string           `json:"message"`
	DeviceList []TrustedDevices `json:"device_list"`
}

type TrustedDevicesResponse struct {
	Message string `json:"message"`
	UserID  string `json:"userid"`
}

type BlockUserResponse struct {
	Message       string `json:"message"`
	BlockedUserID int    `json:"unblocked_user_id"`
}

type UnblockUserResponse struct {
	Message         string `json:"message"`
	UnblockedUserID int    `json:"unblocked_user_id"`
}

type BlockedUsersListResponse struct {
	Message          string             `json:"Message"`
	BlockedUsersList []BlockedUsersList `json:"blocked_users_list"`
	FetchedAt        string             `json:"fetched_at"`
}

type UploadProfileImageResponse struct {
	Message    string `json:"message"`
	UploadedAt string `json:"uploaded_at"`
}

type GetImagesResponse struct {
	Message string  `json:"message"`
	Images  []Image `json:"images"`
}
