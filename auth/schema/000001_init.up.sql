CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE TABLE users
(
    user_id uuid DEFAULT uuid_generate_v4() NOT NULL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    surname VARCHAR(255) NOT NULL,
    username VARCHAR(40) NOT NULL UNIQUE,
    email VARCHAR(255) NOT NULL UNIQUE,
    tel_number VARCHAR(50) NOT NULL UNIQUE,
    id_number VARCHAR(15) NOT NULL UNIQUE,
    password VARCHAR(200) NOT NULL,
    salt VARCHAR(50) NOT NULL,
    date_created VARCHAR(100) NOT NULL,
    ip_address VARCHAR(25) NOT NULL
);
