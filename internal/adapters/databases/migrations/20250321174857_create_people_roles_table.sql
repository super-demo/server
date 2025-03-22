-- +goose Up
-- +goose StatementBegin
CREATE TABLE people_roles (
    people_role_id SERIAL PRIMARY KEY,
    slug VARCHAR(255) NOT NULL,
    description VARCHAR(255),
    site_id INTEGER NOT NULL,
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
DROP TABLE IF EXISTS people_roles;
-- +goose StatementEnd