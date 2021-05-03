-- +goose Up
-- +goose StatementBegin
    CREATE TABLE `client_credentials` (
        `id` bigint unsigned NOT NULL AUTO_INCREMENT,
        `tenant_id` bigint unsigned NOT NULL,
        `name` varchar(255) DEFAULT NULL,
        `code` varchar(100) NOT NULL COMMENT 'AKA client_id',
        `secret` varchar(100) CHARACTER SET utf8 COLLATE utf8_bin  NOT NULL,
        `reference_id` varchar(100) NOT NULL COMMENT 'Reference id of vendor, tenant, user',
        `type` varchar(20) DEFAULT NULL COMMENT 'Possible values are vendor, tenant, user',
        `payload` json DEFAULT NULL COMMENT 'Extra payload used for token and other purpose',
        `is_active` tinyint unsigned DEFAULT '0',
        `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
        `updated_at` timestamp NULL DEFAULT NULL,
        `deleted_at` timestamp NULL DEFAULT NULL,
        PRIMARY KEY (`id`),
        UNIQUE KEY `unique_tenant_id_reference_id_code` (`tenant_id`,`reference_id`,`code`)
    ) ENGINE=InnoDB DEFAULT CHARSET=utf8;
-- +goose StatementEnd

-- +goose StatementBegin
    CREATE TABLE `client_credential_roles` (
        `id` bigint unsigned NOT NULL AUTO_INCREMENT,
        `client_credential_id` bigint unsigned NOT NULL,
        `role_id` bigint unsigned NOT NULL,
        PRIMARY KEY (`id`),
        UNIQUE KEY `unique_client_credential_id_role_id` (`client_credential_id`,`role_id`)
    ) ENGINE=InnoDB DEFAULT CHARSET=utf8;
-- +goose StatementEnd


-- +goose StatementBegin
    ALTER TABLE  `client_credential_roles`
         ADD KEY `client_credential_roles_client_credential_cc_id_idx` (`client_credential_id`),
		 ADD KEY `client_credential_roles_role_id_idx` (`role_id`);
-- +goose StatementEnd

-- +goose StatementBegin
    ALTER TABLE `client_credential_roles`
        ADD  CONSTRAINT `client_credential_roles_client_credential_cc_id_fk` FOREIGN KEY (`client_credential_id`) REFERENCES `client_credentials` (`id`),
        ADD  CONSTRAINT `client_credential_roles_roles_role_id_fk` FOREIGN KEY (`role_id`) REFERENCES `roles` (`id`);
-- +goose StatementEnd


-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
