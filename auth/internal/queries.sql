CREATE TABLE users (
    user_id BIGINT PRIMARY KEY,
    username VARCHAR(50) UNIQUE NOT NULL,
    user_role VARCHAR(15) NOT NULL DEFAULT 'user',
    password VARCHAR(50) NOT NULL,
    first_name VARCHAR(65) NOT NULL,
    last_name VARCHAR(65) NOT NULL,
    country VARCHAR(50) NOT NULL,
    email VARCHAR(100) NOT NULL,
    gender VARCHAR(30) NOT NULL DEFAULT 'Female',
    city  VARCHAR(75) NOT NULL DEFAULT '',
    createdAt VARCHAR(10) NOT NULL DEFAULT '',
    ip_address VARCHAR(30) NOT NULL DEFAULT ''
)