package constants

const (
	RegisterUserQuery = `INSERT INTO "public.User" (personal_number, phonenum, username, email, firstname, lastname, ip_address, password, salt) VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING userid`
	CheckUserQuery    = `SELECT userid FROM "public.User WHERE username=$1 AND password=$2"`
)
