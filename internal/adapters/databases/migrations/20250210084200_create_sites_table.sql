-- +goose Up
-- +goose StatementBegin
CREATE TABLE sites (
    site_id SERIAL PRIMARY KEY,
    site_type_id INT,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    short_description TEXT,
    url TEXT,
    image_url TEXT,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    created_by INTEGER NOT NULL,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_by INTEGER NOT NULL,
    deleted_at TIMESTAMP,

    FOREIGN KEY (site_type_id) REFERENCES site_types(site_type_id),
    FOREIGN KEY (created_by) REFERENCES users(user_id),
    FOREIGN KEY (updated_by) REFERENCES users(user_id)
);

INSERT INTO sites (site_type_id, name, description, short_description, url, image_url, created_by, updated_by) VALUES (2, 'Kasetsart', 'Kasetsart site for test usecases', 'Kasetsart site', 'http://localhost:3001', 'https://localhost:3001', 1, 1);
INSERT INTO sites (site_type_id, name, description, short_description, url, image_url, created_by, updated_by) VALUES (1, 'Site 1', 'Site 1 for test usecases', 'Site 1', 'http://localhost:3001', 'https://localhost:3001', 1, 1);
INSERT INTO sites (site_type_id, name, description, short_description, url, image_url, created_by, updated_by) VALUES (1, 'Site 2', 'Site 2 for test usecases', 'Site 2', 'http://localhost:3001', 'https://localhost:3001', 1, 1);
INSERT INTO sites (site_type_id, name, description, short_description, url, image_url, created_by, updated_by) VALUES (1, 'Site 3', 'Site 3 for test usecases', 'Site 3', 'http://localhost:3001', 'https://localhost:3001', 1, 1);
INSERT INTO sites (site_type_id, name, description, short_description, url, image_url, created_by, updated_by) VALUES (1, 'Site 4', 'Site 4 for test usecases', 'Site 4', 'http://localhost:3001', 'https://localhost:3001', 1, 1);
INSERT INTO sites (site_type_id, name, description, short_description, url, image_url, created_by, updated_by) VALUES (1, 'Site 5', 'Site 5 for test usecases', 'Site 5', 'http://localhost:3001', 'https://localhost:3001', 1, 1);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS sites;
-- +goose StatementEnd
