package constants

const (
	RegisterUserQuery                  = `INSERT INTO users (personal_number, phonenum, username, email, firstname, lastname, ip_address, password, salt) VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING userid`
	CheckUserQuery                     = `SELECT userid FROM users WHERE username=$1 AND password=$2`
	CheckUserByEmail                   = `SELECT userid FROM users WHERE email=$1 and username=$2 and personal_number=$3`
	InsertProfileActivityResetPassword = `UPDATE profileactivities set userid = (
												select userid from users where personal_number = $1
												),
												last_reset_password = $2,
												last_password_change = $3,
												location = $4`
	InsertProfileActivityResetPasswordProfile = `UPDATE profileactivities set userid = (
													select userid from users where username = $1
													),
													last_reset_password = $2,
													last_password_change = $3,
													location = $4`
	UpdatePassword            = `UPDATE users set password = $1 where personal_number = $2`
	UpdatePasswordFromProfile = `Update users set password = $1 where username = $2`
	GetProfileDetails         = `SELECT * FROM userinfo where userid = $1`
	GetUserInfo               = `SELECT * FROM users
									 INNER JOIN
									 userinfo u ON users.userid = u.userid
									 WHERE users.userid = $1`
	AddTrustedDevice = `INSERT INTO trusteddevices (userid, device_id, device_name, device_ip_address, device_identifier)
								   VALUES($1, $2, $3 , $4, $5)`
	GetTrustedDevices = `select userid, device_id, device_name, device_ip_address,device_identifier
    							   from trusteddevices where userid = $1`
	BlockUserQuery = `INSERT INTO blockedusers (userid, blocked_user_id, blocked_at) VALUES
						(
							$1, $2, $3
						)`
	UnblockUserQuery = `DELETE  from blockedusers
    							   where blocked_user_id = $1 and userid = $2`
	GetBlockedUsersList = `SELECT * FROM blockedusers WHERE userid = $1`
	AddProfileImage     = `INSERT INTO image (userid, profileimage, uploadedat) VALUES ($1, $2, $3)`
)
