CREATE TABLE users (
    ID SERIAL PRIMARY KEY,
    full_name VARCHAR NOT NULL,
    email VARCHAR UNIQUE NOT NULL,
    username VARCHAR UNIQUE NOT NULL,
    password VARCHAR NOT NULL,
    deposit FLOAT NOT NULL DEFAULT 0
);