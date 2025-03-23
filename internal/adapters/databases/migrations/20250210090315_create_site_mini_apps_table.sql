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

-- INSERT INTO site_mini_apps (site_id, slug, description, link_url, image_url, created_by, updated_by) VALUES (1, 'Ku Book', 'Mini app 1 for test usecases', 'http://localhost:3002', 'https://cdn.jsdelivr.net/npm/@fortawesome/fontawesome-free@6.4.2/svgs/solid/book.svg', 1, 1);
-- INSERT INTO site_mini_apps (site_id, slug, description, link_url, image_url, created_by, updated_by) VALUES (1, 'Ku Research', 'Mini app 2 for test usecases', 'http://localhost:3003', 'https://cdn.jsdelivr.net/npm/@fortawesome/fontawesome-free@6.4.2/svgs/regular/file.svg', 1, 1);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS site_mini_apps;
-- +goose StatementEnd
