package constants

/* Success Messages */

const (
	CreatedUserSuccess  = "[Register Handler] - User Created Successfully"
	LoggedInUserSuccess = "[Login Handler] - User Logged In Successfully"
)

/* Error Messages */

const (
	UserAlreadyRegisteredError = "[Register Handler] - User Already Registered"
	UserNotFoundError          = "[Login Handler] - User With These Credentials Not Found"

	InvalidTokenError  = "[Refresh Login Handler] - Invalid Refresh Token"
	InvalidTokenClaims = "[Token Handler] - invalid Token Claims"
)

/* Generic Errors */

const (
	InternalServerError = "[Generic Handler] - Internal Server Error"
	BadRequest          = "[Generic Handler] - Some Of Your Credentials Are Missing"
)

/* Keys */

const (
	RedisSalt = "Salt"
)
