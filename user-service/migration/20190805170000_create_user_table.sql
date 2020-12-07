-- +goose Up
-- SQL in this section is executed when the migration is applied.
CREATE TABLE IF NOT EXISTS users
(
    id             INT UNSIGNED NOT NULL AUTO_INCREMENT,
    name           VARCHAR(255) NOT NULL,
    email          VARCHAR(255) NOT NULL,
    created_date   DATE         NOT NULL,
    first_name     VARCHAR(255) NULL,
    created_at     TIMESTAMP    NOT NULL,
    updated_at     TIMESTAMP    NULL,
    deleted_at     TIMESTAMP    NULL,
    PRIMARY KEY (id)
);

