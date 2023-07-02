--migrate create --ext sql -dir ./schema -seq init
--migrate -path ./schema -database 'postgres://crudapp:crudapp@localhost:5432/crudapp?sslmode=disable' up
CREATE TABLE books
(
    id SERIAL NOT NULL UNIQUE,
    title VARCHAR(255) NOT NULL,
    author VARCHAR NOT NULL,
    publish_date TIMESTAMP NOT NULL DEFAULT (now()),
    rating VARCHAR NOT NULL
);

CREATE TABLE users
(
    id SERIAL NOT NULL UNIQUE,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL DEFAULT '',
    registered_at TIMESTAMP NOT NULL DEFAULT (now())
);