UPDATE profileactivities set
            userid = (
                select userid from users where personal_number = '01027090407' and username = 'davit.darsalia' and email = 'darsalia.david1998@gmail.com'
            ),
            last_reset_password = '1998-08-08',
            last_password_change = '1998-08-08',
            location = 'Tbilisi'
