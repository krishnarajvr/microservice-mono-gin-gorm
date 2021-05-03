-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS tenants
(
    id             BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
    name           VARCHAR(255) NOT NULL,
    code           VARCHAR(50)  NOT NULL,
    domain         VARCHAR(20)  NOT NULL,
    secret         VARCHAR(100) NOT NULL COMMENT 'Used for tenant token',
    email          VARCHAR(255) NOT NULL,
    is_active      tinyint(1)   unsigned DEFAULT '0',
    created_at     TIMESTAMP    NOT NULL,
    updated_at     TIMESTAMP    NULL,
    deleted_at     TIMESTAMP    NULL,
    PRIMARY KEY (id),
    UNIQUE KEY `unique_code` (`code`),
    UNIQUE KEY `unique_name` (`name`),
    UNIQUE KEY `unique_email` (`email`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8;
-- +goose StatementEnd

-- +goose StatementBegin
CREATE TABLE tenant_settings (
  id             BIGINT UNSIGNED NOT NULL,
  tenant_id      BIGINT UNSIGNED NOT NULL DEFAULT 0,
  settings       JSON NOT NULL,
  type           VARCHAR(20)  NOT NULL DEFAULT 'SYSTEM' COMMENT 'Settings type SYSTEM, ACCOUNT, NOTIFICATION etc',
  level          TINYINT NOT NULL DEFAULT 0 COMMENT 'Level of the seetings 0: Core, 1: Domain, 2: Instance',
  created_at     TIMESTAMP    NOT NULL,
  updated_at     TIMESTAMP    NULL,
  deleted_at     TIMESTAMP    NULL,
  PRIMARY KEY (id),
  UNIQUE KEY `unique_tenant_id_type_level` (`tenant_id`,`type`,`level`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
