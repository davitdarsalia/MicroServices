package queries

const (
	RegisterUserQuery = `INSERT INTO microservices.public.users (username, user_role, password, first_name, last_name, country, email, gender, city, createdat, ip_address) VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11) RETURNING user_id`
	CheckUserQuery    = `SELECT user_id FROM microservices.public.users WHERE email = $1 AND password = $2`
	CheckUserForReset = `SELECT user_id FROM microservices.public.users WHERE email = $1 AND id_number = $2`
	ResetPassword     = `UPDATE microservices.public.users SET password = $1 WHERE user_id = $2`
)
