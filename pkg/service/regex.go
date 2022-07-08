package service

const (
	passwordRegex       = `^[A-Z][a-zA-Z_1234567890.,/?!@#$%^&*()";:<>,./]{8,}$`
	nameRegex           = `[A-Z]{1,}[a-zA-Z\s]{2,}[A-Z]{1,}[a-z]{1,}[^!@#$%^&*~./;?<>"=|]`
	emailRegex          = `[a-z_.123456789]{3,}[^!@#$%^&*()-+=|}{;:",./]@gmail.com`
	usernameRegex       = `^[a-zA-Z0-9_@.!-#$%^&*]{8,}`
	ipAddressRegex      = `[0-9]{1,3}.[0-9]{1,3}.[0-9]{1,3}.[0-9]{1,3}`
	mobileNumberRegex   = `[0-9_+]{10,}[^!@#$%^&*()_+<>,.'?/:;"'[\]]`
	personalNumberRegex = `[0-9]{10,15}`
)
