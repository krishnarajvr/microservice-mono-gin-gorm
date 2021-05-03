-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS roles
(
    id             BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
    tenant_id      BIGINT UNSIGNED NOT NULL DEFAULT 0,
    name           VARCHAR(255)    NOT NULL,
    code           VARCHAR(255)    NOT NULL,
    created_at     TIMESTAMP       NOT NULL,
    updated_at     TIMESTAMP       NULL,
    deleted_at     TIMESTAMP       NULL,
    PRIMARY KEY (id),
    UNIQUE KEY `unique_tenant_id_name` (`tenant_id`,`name`),
    UNIQUE KEY `unique_tenant_id_code` (`tenant_id`,`code`)
);
-- +goose StatementEnd
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS users
(
    id              BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
    tenant_id       BIGINT UNSIGNED NOT NULL,
    name            VARCHAR(255)    NOT NULL,
    email           VARCHAR(255)    NOT NULL,
    first_name      VARCHAR(255)    NULL,
    last_name       VARCHAR(255)    NULL,
    is_active       TINYINT(1) UNSIGNED DEFAULT 0 ,
    is_root         TINYINT(1) UNSIGNED DEFAULT 0 COMMENT 'First User, Cannot delete',
    password        VARCHAR(255) DEFAULT NULL,
    login_at        TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    login_count     INT DEFAULT 0,
    created_user_id BIGINT UNSIGNED NOT NULL DEFAULT 0,
    created_at      TIMESTAMP       NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at      TIMESTAMP       NULL,
    deleted_at      TIMESTAMP       NULL,
    PRIMARY KEY (id),
    UNIQUE KEY `unique_tenant_id_name` (`tenant_id`,`name`),
    UNIQUE KEY `unique_tenant_id_email` (`tenant_id`,`email`)
);
-- +goose StatementEnd
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS user_audit_log
(
   id             BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
   tenant_id      BIGINT UNSIGNED NOT NULL,
   user_id        BIGINT UNSIGNED NOT NULL,
   audit_data     JSON NOT NULL,
   created_at     TIMESTAMP   NOT NULL  DEFAULT CURRENT_TIMESTAMP,
   PRIMARY KEY (id)
)
-- +goose StatementEnd
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS user_roles
(
    id             BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
    user_id        BIGINT UNSIGNED NOT NULL,
    role_id        BIGINT UNSIGNED NOT NULL,
    PRIMARY KEY (id),
    UNIQUE KEY `unique_user_id_role_id` (`user_id`,`role_id`)
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
