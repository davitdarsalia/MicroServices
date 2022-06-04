UPDATE profileactivities set
            userid = (
                select userid from users where personal_number = $1 and username = $2 and email = $3
            ),
            last_reset_password = '1998-08-08',
            last_password_change = '1998-08-08',
            location = 'Tbilisi';

UPDATE trusted_devices set
            device_id = 1,
            userid = 23
    where userid = 23;


CREATE TABLE userinfo (
            profileImage  bytea,
            followers INT DEFAULT 0 NOT NULL,
            following INT DEFAULT 0 NOT NULL,
            blocked_users_amount INT DEFAULT 0 NOT NULL,
            working_place VARCHAR(200),
            education VARCHAR(200),
            origin VARCHAR(200),
            additional_email VARCHAR(200),
            userid BIGINT REFERENCES users (userid)
)

