CREATE TABLE books
(
    id SERIAL NOT NULL UNIQUE,
    title VARCHAR(255) NOT NULL,
    author VARCHAR NOT NULL,
    publish_date TIMESTAMP NOT NULL DEFAULT (now()),
    rating VARCHAR NOT NULL
);