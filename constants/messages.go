package constants

// Errors
const (
	UserNotFoundError          = "[LogIn Handler] - User Not Found"
	BadRequest                 = "[All Handlers] - Some Of Your Credentials Are Missing"
	InternalServerError        = "[All Handlers] - Internal Server Error"
	UserAlreadyRegistered      = "[Register Handler] - User Already Registered"
	ResetPasswordError         = "[Reset Password Handler] - Some Of Your Credentials Are Invalid\n Failed To Reset Password"
	ValidateResetPasswordError = "[Validate Password Handler] - Incorrect Validation Code, Try Again "
	GetProfileDetailsError     = "[Get Profile Details Handler] - Failed To Fetch Profile Details"
	GetUserInfoError           = "[Get User Info Handler] - Failed To Fetch UserInfo"
	UpdateTrustedDevicesError  = "[Update Trusted Devices Handler] - Failed To Add A Trusted Device"
)

// Success
const (
	CreatedUserSuccess           = "[Register Handler] - User Created Successfully"
	SuccessfulSignIn             = "[LogIn Handler] - User Successfully Signed In"
	ResetPasswordSuccess         = "[Reset Password Handler] - Password Email Sent Successfully"
	ValidateResetPasswordSuccess = "[Validate Password Handler] - Password Changed Successfully"
	GetProfileDetailsSuccess     = "[Validate Password Handler] - Profile Details Fetched Successfully"
	GetUserInfoSuccess           = "[Get User Info Handler] - Failed To Fetch UserInfo"
	UpdateTrustedDevicesSuccess  = "[Update Trusted Devices Handler] - Trusted Device Added Successfully"
)
