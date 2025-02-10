-- +goose Up
-- +goose StatementBegin
CREATE TABLE site_users (
    site_user_id SERIAL PRIMARY KEY,
    site_id INTEGER NOT NULL,
    user_id INTEGER NOT NULL,
    user_level_id INTEGER NOT NULL,
    is_active BOOLEAN NOT NULL DEFAULT TRUE,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    created_by INTEGER NOT NULL,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_by INTEGER NOT NULL,
    deleted_at TIMESTAMP,

    FOREIGN KEY (site_id) REFERENCES sites(site_id),
    FOREIGN KEY (user_id) REFERENCES users(user_id),
    FOREIGN KEY (user_level_id) REFERENCES user_levels(user_level_id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS site_users;
-- +goose StatementEnd
