-- +goose Up
-- +goose StatementBegin
CREATE TABLE site_peoples (
    site_people_id SERIAL PRIMARY KEY,
    site_id INTEGER NOT NULL,
    user_id INTEGER NOT NULL,
    is_active BOOLEAN NOT NULL DEFAULT TRUE,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    created_by INTEGER NOT NULL,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_by INTEGER NOT NULL,
    deleted_at TIMESTAMP,

    FOREIGN KEY (site_id) REFERENCES sites(site_id),
    FOREIGN KEY (user_id) REFERENCES users(user_id),
    FOREIGN KEY (created_by) REFERENCES users(user_id),
    FOREIGN KEY (updated_by) REFERENCES users(user_id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS site_peoples;
-- +goose StatementEnd
