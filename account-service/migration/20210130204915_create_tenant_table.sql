-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS tenants
(
    id             INT UNSIGNED NOT NULL AUTO_INCREMENT,
    name           VARCHAR(255) NOT NULL,
    code           VARCHAR(50)  NOT NULL,
    domain         VARCHAR(20)  NOT NULL,
    secret         VARCHAR(100) NOT NULL,
    email          VARCHAR(255) NOT NULL,
    is_active      bit(1)       DEFAULT 0,
    created_at     TIMESTAMP    NOT NULL,
    updated_at     TIMESTAMP    NULL,
    deleted_at     TIMESTAMP    NULL,
    PRIMARY KEY (id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
