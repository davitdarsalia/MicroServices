package queries

const CreateUserQuery = `INSERT INTO users (name, surname, username, email, tel_number, id_number, password, date_created, ip_address, salt) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10) RETURNING user_id`
const LoginUserQuery = `SELECT user_id, salt, password from users where email = $1 and id_number = $2`
const UpdatePasswordQuery = `UPDATE users SET password = $1 WHERE email = $2 and id_number = $3 and tel_number = $4`

const RecoverPasswordFirstStepQuery = `
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

DO $$
    DECLARE
        id uuid;
    BEGIN
        SELECT user_id INTO id
        FROM (
                 SELECT user_id  FROM users WHERE
                         email = $1 and
                         id_number = $2 and
                         tel_number = $3
             ) AS subquery;

        IF id IS NOT NULL  AND id::text ~ '^[a-f0-9]{8}-[a-f0-9]{4}-4[a-f0-9]{3}-[8|9|aA|bB][a-f0-9]{3}-[a-f0-9]{12}$'
            THEN
                UPDATE users SET password = '221111222' WHERE user_id = id;
        END IF;
END;$$;
`
