-- +goose Up
-- +goose StatementBegin
CREATE TABLE site_users (
    site_user_id SERIAL PRIMARY KEY,
    site_id INTEGER NOT NULL,
    user_id INTEGER NOT NULL,
    site_user_level_id INTEGER NOT NULL,
    is_active BOOLEAN NOT NULL DEFAULT TRUE,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    created_by INTEGER NOT NULL,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_by INTEGER NOT NULL,
    deleted_at TIMESTAMP,

    FOREIGN KEY (site_id) REFERENCES sites(site_id),
    FOREIGN KEY (user_id) REFERENCES users(user_id),
    FOREIGN KEY (site_user_level_id) REFERENCES user_levels(user_level_id),
    FOREIGN KEY (created_by) REFERENCES users(user_id),
    FOREIGN KEY (updated_by) REFERENCES users(user_id)
);

INSERT INTO site_users (site_id, user_id, site_user_level_id, created_by, updated_by) VALUES (1, 1, 1, 1, 1);
INSERT INTO site_users (site_id, user_id, site_user_level_id, created_by, updated_by) VALUES (1, 2, 1, 1, 1);
INSERT INTO site_users (site_id, user_id, site_user_level_id, created_by, updated_by) VALUES (1, 3, 5, 1, 1);

INSERT INTO site_users (site_id, user_id, site_user_level_id, created_by, updated_by) VALUES (2, 1, 2, 1, 1);
INSERT INTO site_users (site_id, user_id, site_user_level_id, created_by, updated_by) VALUES (2, 2, 2, 1, 1);

INSERT INTO site_users (site_id, user_id, site_user_level_id, created_by, updated_by) VALUES (3, 1, 3, 1, 1);
INSERT INTO site_users (site_id, user_id, site_user_level_id, created_by, updated_by) VALUES (3, 2, 3, 1, 1);

INSERT INTO site_users (site_id, user_id, site_user_level_id, created_by, updated_by) VALUES (4, 1, 4, 1, 1);
INSERT INTO site_users (site_id, user_id, site_user_level_id, created_by, updated_by) VALUES (4, 2, 4, 1, 1);

INSERT INTO site_users (site_id, user_id, site_user_level_id, created_by, updated_by) VALUES (5, 1, 5, 1, 1);
INSERT INTO site_users (site_id, user_id, site_user_level_id, created_by, updated_by) VALUES (5, 2, 5, 1, 1);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS site_users;
-- +goose StatementEnd
