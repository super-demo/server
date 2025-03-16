-- +goose Up
-- +goose StatementBegin
CREATE TABLE site_mini_apps (
    site_mini_app_id SERIAL PRIMARY KEY,
    site_id INTEGER NOT NULL,
    slug VARCHAR(255) NOT NULL,
    description TEXT,
    link_url TEXT,
    image_url TEXT,
    is_active BOOLEAN NOT NULL DEFAULT TRUE,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    created_by INTEGER NOT NULL,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_by INTEGER NOT NULL,
    deleted_at TIMESTAMP,

    FOREIGN KEY (site_id) REFERENCES sites(site_id),
    FOREIGN KEY (created_by) REFERENCES users(user_id),
    FOREIGN KEY (updated_by) REFERENCES users(user_id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS site_mini_apps CASCADE;
-- +goose StatementEnd
