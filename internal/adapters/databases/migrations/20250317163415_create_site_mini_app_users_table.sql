-- +goose Up
-- +goose StatementBegin
CREATE TABLE site_mini_app_users (
    site_mini_app_user_id SERIAL PRIMARY KEY,
    site_mini_app_id INTEGER NOT NULL,
    user_id INTEGER NOT NULL,
    is_active BOOLEAN NOT NULL DEFAULT TRUE,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    created_by INTEGER NOT NULL,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,            
    updated_by INTEGER NOT NULL,
    deleted_at TIMESTAMP,

    FOREIGN KEY (site_mini_app_id) REFERENCES site_mini_apps(site_mini_app_id),
    FOREIGN KEY (user_id) REFERENCES users(user_id),
    FOREIGN KEY (created_by) REFERENCES users(user_id),
    FOREIGN KEY (updated_by) REFERENCES users(user_id)
);

-- INSERT INTO site_mini_app_users (site_mini_app_id, user_id, created_by, updated_by) VALUES (1, 1, 1, 1);
-- INSERT INTO site_mini_app_users (site_mini_app_id, user_id, created_by, updated_by) VALUES (1, 2, 1, 1);
-- INSERT INTO site_mini_app_users (site_mini_app_id, user_id, created_by, updated_by) VALUES (1, 3, 1, 1);

-- INSERT INTO site_mini_app_users (site_mini_app_id, user_id, created_by, updated_by) VALUES (2, 1, 1, 1);
-- INSERT INTO site_mini_app_users (site_mini_app_id, user_id, created_by, updated_by) VALUES (2, 2, 1, 1);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS site_mini_app_users;
-- +goose StatementEnd
