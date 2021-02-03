-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS tenants
(
    id             BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
    name           VARCHAR(255) NOT NULL,
    code           VARCHAR(50)  NOT NULL,
    domain         VARCHAR(20)  NOT NULL,
    secret         VARCHAR(100) NOT NULL,
    email          VARCHAR(255) NOT NULL,
    is_active      tinyint(1)   unsigned DEFAULT '0',
    created_at     TIMESTAMP    NOT NULL,
    updated_at     TIMESTAMP    NULL,
    deleted_at     TIMESTAMP    NULL,
    PRIMARY KEY (id),
    UNIQUE KEY `code` (`code`),
    UNIQUE KEY `name` (`name`),
    UNIQUE KEY `email` (`email`)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
