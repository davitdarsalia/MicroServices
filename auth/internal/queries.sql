-- Users Table
CREATE TABLE users (
    user_id SERIAL NOT NULL UNIQUE PRIMARY KEY ,
    username VARCHAR(50) UNIQUE NOT NULL,
    user_role VARCHAR(15) NOT NULL DEFAULT 'user',
    password VARCHAR(50) NOT NULL,
    first_name VARCHAR(65) NOT NULL,
    last_name VARCHAR(65) NOT NULL,
    country VARCHAR(50) NOT NULL,
    email VARCHAR(100) NOT NULL,
    gender VARCHAR(30) NOT NULL DEFAULT 'Female',
    city  VARCHAR(75) NOT NULL DEFAULT '',
    createdAt VARCHAR(30) NOT NULL DEFAULT '',
    ip_address VARCHAR(30) NOT NULL DEFAULT '',
    verified BOOL NOT NULL DEFAULT false,
    id_number VARCHAR(25) NOT NULL UNIQUE DEFAULT ''
);

-- Inserting User Sample Data Into Users Table
INSERT INTO users (username, user_role, password, first_name, last_name, country, email , gender, city, createdat,ip_address)
VALUES ('Davit1998', 'Admin', 'David.19982','David', 'Darsalia','Georgia', 'darsalia.david1998@gmail.com', 'Male', 'Tbilisi', '22.01.2022', '01.1.1.1' ) RETURNING user_id;

-- Delete All Values
DELETE  FROM users WHERE user_id > 1