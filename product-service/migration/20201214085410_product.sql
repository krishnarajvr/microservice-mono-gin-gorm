-- +goose Up
-- SQL in this section is executed when the migration is applied.
CREATE TABLE IF NOT EXISTS products
(
    id             INT UNSIGNED NOT NULL AUTO_INCREMENT,
    name           VARCHAR(255) NOT NULL,
    code           VARCHAR(100) NOT NULL,
    description    TEXT         NULL,
    created_at     TIMESTAMP    NOT NULL,
    updated_at     TIMESTAMP    NULL,
    deleted_at     TIMESTAMP    NULL,
    PRIMARY KEY (id)
);

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
