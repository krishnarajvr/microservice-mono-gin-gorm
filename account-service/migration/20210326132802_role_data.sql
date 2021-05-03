-- +goose Up
-- +goose StatementBegin
ALTER TABLE `roles` 
    ADD `level` SMALLINT NOT NULL AFTER `code`; 
-- +goose StatementEnd

-- +goose StatementBegin
INSERT INTO `roles` 
(
    `id`, 
    `tenant_id`, 
    `name`, 
    `code`, 
    `level`, 
    `created_at`, 
    `updated_at`, 
    `deleted_at`
) VALUES
(1, 0, 'Super Admin', 'SUPER_ADMIN', 0, '2021-03-26 07:54:55', NULL, NULL),
(2, 0, 'Tenant Admin', 'ADMIN', 1, '2021-03-26 07:54:55', NULL, NULL),
(3, 0, 'Tenant Admin View', 'ADMIN_VIEW', 1, '2021-03-26 08:00:52', NULL, NULL),
(4, 0, 'Vendor', 'VENDOR', 2, '2021-03-26 08:00:09', NULL, NULL),
(5, 0, 'End User', 'USER', 3, '2021-03-26 08:02:43', NULL, NULL);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
