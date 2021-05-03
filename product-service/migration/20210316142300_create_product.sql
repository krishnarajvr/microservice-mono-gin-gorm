-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS products
(
    id             BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
    tenant_id      BIGINT UNSIGNED NOT NULL,
    name           VARCHAR(255) NOT NULL,
    code           VARCHAR(50)  NOT NULL,
    description    TEXT DEFAULT NULL,
    meta           JSON DEFAULT NULL,
    is_active      tinyint(1)   unsigned DEFAULT '0',
    created_at     TIMESTAMP    NOT NULL,
    updated_at     TIMESTAMP    NULL,
    deleted_at     TIMESTAMP    NULL,
    PRIMARY KEY (id),
    UNIQUE KEY `unique_tenant_id_name` (`tenant_id`,`name`),
    UNIQUE KEY `unique_tenant_id_code` (`tenant_id`,`code`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8;
-- +goose StatementEnd


-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
