-- +goose Up
-- +goose StatementBegin
CREATE TABLE organization_users (
    organization_user_id SERIAL PRIMARY KEY,
    organization_id INTEGER NOT NULL,
    user_id INTEGER NOT NULL,
    user_level_id INTEGER NOT NULL,
    is_active BOOLEAN NOT NULL DEFAULT TRUE,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    created_by INTEGER NOT NULL,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_by INTEGER NOT NULL,
    deleted_at TIMESTAMP,

    FOREIGN KEY (organization_id) REFERENCES organizations(organization_id),
    FOREIGN KEY (user_id) REFERENCES users(user_id),
    FOREIGN KEY (user_level_id) REFERENCES user_levels(user_level_id),
    FOREIGN KEY (created_by) REFERENCES users(user_id),
    FOREIGN KEY (updated_by) REFERENCES users(user_id)
)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS organization_users;
-- +goose StatementEnd
