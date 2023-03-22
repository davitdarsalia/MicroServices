package responses

/* Success */

const (
	CreateUserSuccessMessage        = `[Create User Handler] - User Created Successfully`
	LoggedInUserSuccessMessage      = `[Log In User Handler] - User Logged In Successfully`
	RecoveredPasswordSuccessMessage = `[Recovered Password Handler] - Code Sent To Your Email`
)

/* Errors */

const (
	CreateUserErrorMessage = `[Create User Handler] - User Already Exists`
	LogInUserErrorMessage  = `[Log In User Handler] - User Not Found`
	BadRequestErrorMessage = `[Generic Message] - Some Of Your Credentials Are Missing`
)

/* Constants */

const (
	ValidationFailedErrorMessage = `verifications failed for fields`
)
