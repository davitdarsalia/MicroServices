UPDATE profileactivities set
            userid = (
                select userid from users where personal_number = $1 and username = $2 and email = $3
            ),
            last_reset_password = $4,
            last_password_change = $5,
            location = $6
