-- +goose Up
-- +goose StatementBegin
CREATE TABLE organization_category_users (
    organization_category_user_id SERIAL PRIMARY KEY,
    organization_category_id INTEGER NOT NULL,
    user_id INTEGER NOT NULL,
    user_level_id INTEGER NOT NULL,
    is_active BOOLEAN NOT NULL DEFAULT TRUE,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    created_by INTEGER NOT NULL,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_by INTEGER NOT NULL,
    deleted_at TIMESTAMP,

    FOREIGN KEY (organization_category_id) REFERENCES organization_categories(organization_category_id),
    FOREIGN KEY (user_id) REFERENCES users(user_id),
    FOREIGN KEY (user_level_id) REFERENCES user_levels(user_level_id),
    FOREIGN KEY (created_by) REFERENCES users(user_id),
    FOREIGN KEY (updated_by) REFERENCES users(user_id)
)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS organization_category_users;
-- +goose StatementEnd
