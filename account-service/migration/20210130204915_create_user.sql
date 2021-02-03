-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS users
(
    id             BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
    tenant_id      BIGINT UNSIGNED NOT NULL,
    name           VARCHAR(255)    NOT NULL,
    email          VARCHAR(255)    NOT NULL,
    first_name     VARCHAR(255)    NULL,
    last_name      VARCHAR(255)    NULL,
    is_active      TINYINT(1) UNSIGNED DEFAULT '0',
    created_at     TIMESTAMP       NOT NULL,
    updated_at     TIMESTAMP       NULL,
    deleted_at     TIMESTAMP       NULL,
    PRIMARY KEY (id),
    UNIQUE KEY `name` (`name`),
    UNIQUE KEY `email` (`email`)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
