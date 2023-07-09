CREATE TABLE refresh_tokens
(
    id SERIAL NOT NULL UNIQUE,
    user_id SERIAL NOT NULL UNIQUE,
    token VARCHAR(255) NOT NULL UNIQUE,
    expires_at TIMESTAMP NOT NULL DEFAULT (now())
);