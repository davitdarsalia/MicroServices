package constants

// Errors
const (
	InvalidAuthError             = "[Auth Checking Middleware] - Invalid Auth Internals"
	UserNotFoundError            = "[LogIn Handler] - User Not Found"
	BadRequest                   = "[All Handlers] - Some Of Your Credentials Are Missing"
	InternalServerError          = "[All Handlers] - Internal Server Error"
	UserAlreadyRegistered        = "[Register Handler] - User Already Registered"
	ResetPasswordError           = "[Reset Password Handler] - Some Of Your Credentials Are Invalid\n Failed To Reset Password"
	ValidateResetPasswordError   = "[Validate Password Handler] - Incorrect Validation Code, Try Again "
	GetProfileDetailsError       = "[Get Profile Details Handler] - Failed To Fetch Profile Details"
	GetUserInfoError             = "[Get User Info Handler] - Failed To Fetch UserInfo"
	AddTrustedDeviceError        = "[Add Trusted Device Handler] - Failed To Add A Trusted Device"
	GetTrustedDevicesError       = "[Get Trusted Devices Handler] - Failed To Fetch Trusted Devices List"
	BlockUserError               = "[Block User Handler] - Failed To Block User"
	UnblockUserError             = "[Unblock User Handler] - Failed To Unblock User"
	GetBlockedUserListError      = "[Blocked Users List Handler] - Failed To Fetch Blocked Users List"
	GetSessionIDError            = "[Session Manager || Middleware ||] - Unauthorized"
	ProfileImageWithoutFileError = "[Profile Image Handler] - File Is Missing. Please, Upload  A File"
	GetImageErrors               = "[Image Getter Handler] - Failed To Fetch User Images"
)

// Success
const (
	CreatedUserSuccess           = "[Register Handler] - User Created Successfully"
	SuccessfulSignIn             = "[LogIn Handler] - User Successfully Signed In"
	ResetPasswordSuccess         = "[Reset Password Handler] - Password Email Sent Successfully"
	ValidateResetPasswordSuccess = "[Validate Password Handler] - Password Changed Successfully"
	GetProfileDetailsSuccess     = "[Validate Password Handler] - Profile Details Fetched Successfully"
	GetUserInfoSuccess           = "[Get User Info Handler] - UserInfo Fetched Successfully"
	AddTrustedDeviceSuccess      = "[Add Trusted Device Handler] - Trusted Device Added Successfully"
	GetTrustedDevicesSuccess     = "[Get Trusted Devices Handler] - UserInfo Fetched Successfully"
	BlockUserSuccess             = "[Block User Handler] - User Blocked Successfully"
	UnblockUserSuccess           = "[Unblock User Handler] - User Unblocked Successfully"
	GetBlockedUserListSuccess    = "[Blocked Users List Handler] - Blocked Users List Fetched Successfully"
	ProfileImageUploadSuccess    = "[Profile Image Handler] - Profile Image Uploaded Successfully"
	GetImageSuccess              = "[Image Getter Handler] - Images Fetched Successfully"
)
