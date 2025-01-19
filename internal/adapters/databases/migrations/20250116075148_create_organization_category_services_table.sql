-- +goose Up
-- +goose StatementBegin
CREATE TABLE organization_category_services (
    organization_category_service_id SERIAL PRIMARY KEY,
    organization_category_id INTEGER NOT NULL,
    organization_service_id INTEGER NOT NULL,
    organization_id INTEGER NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    created_by INTEGER NOT NULL,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_by INTEGER NOT NULL,
    deleted_at TIMESTAMP,

    FOREIGN KEY (organization_category_id) REFERENCES organization_categories(organization_category_id),
    FOREIGN KEY (organization_service_id) REFERENCES organization_services(organization_service_id),
    FOREIGN KEY (created_by) REFERENCES users(user_id),
    FOREIGN KEY (updated_by) REFERENCES users(user_id)
)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS organization_category_services;
-- +goose StatementEnd
