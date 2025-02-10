-- +goose Up
-- +goose StatementBegin
CREATE TABLE sites (
    site_id SERIAL PRIMARY KEY,
    site_type_id INT,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    created_by INTEGER NOT NULL,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_by INTEGER NOT NULL,
    deleted_at TIMESTAMP,

    FOREIGN KEY (site_type_id) REFERENCES site_types(site_type_id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS sites;
-- +goose StatementEnd
