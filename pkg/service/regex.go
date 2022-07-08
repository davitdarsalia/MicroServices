package service

const (
	EmailRegex     = `\w+([-+.]\w+)*@(?!(hotmail|gmail|yahoo|msn|excite|lycos|aol|live|mail)\.com$)\w+([-.]\w+)*\.\w+([-.]\w+)*`
	UsernameRegex  = `^[a-zA-Z0-9][a-zA-Z0-9_]{2,29}$`
	NameRegex      = `([a-zA-Z]{1,64}\s*)+`
	PasswordRegex  = `^(?=.*\d)(?=.*[a-z])(?=.*[A-Z])(?=.*[a-zA-Z]).{8,}$`
	IPAddressRegex = `^((25[0-5]|(2[0-4]|1\d|[1-9]|)\d)(\.(?!$)|$)){4}$`
)
