package queries

const CreateUserQuery = `INSERT INTO users (name, surname, username, email, tel_number, id_number, password, date_created, ip_address, salt) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10) RETURNING user_id`
const LoginUserQuery = `SELECT user_id, salt, password from users where email = $1 and id_number = $2`
const UpdatePasswordQuery = `UPDATE users SET password = $1 WHERE email = $2 and id_number = $3 and tel_number = $4`
