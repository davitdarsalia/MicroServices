package queries

const RegisterUserQuery = `INSERT INTO microservices.public.users (username, user_role, password, first_name, last_name, country, email, gender, city, createdat, ip_address) VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11) RETURNING user_id`
