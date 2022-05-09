package constants

const (
	REGISTER_USER_QUERY = `INSERT INTO "public.User" (personal_number, phonenum, username, email, firstname, lastname, ip_address, password, salt) VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING userid`
)
