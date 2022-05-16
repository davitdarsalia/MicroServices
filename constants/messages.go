package constants

// Errors
const (
	UserNotFoundError     = "[LogIn Handler] - User Not Found"
	BadRequest            = "[All Handlers] - Some Of Your Credentials Are Missing"
	UserAlreadyRegistered = "[Register Handler] - User Already Registered"
	ResetPasswordError    = "[Reset Password Handler] - Some Of Your Credentials Are Invalid\n Failed To Reset Password"
)

// Success
const (
	CreatedUserSuccess   = "[Register Handler] - User Created Successfully"
	SuccessfulSignIn     = "[LogIn Handler] - User Successfully Signed In"
	ResetPasswordSuccess = "[Reset Password Handler] - Password Reset Successfully"
)
