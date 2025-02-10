-- +goose Up
-- +goose StatementBegin
CREATE TABLE site_types (
  site_type_id SERIAL PRIMARY KEY,
  slug VARCHAR(255) NOT NULL,
  description VARCHAR(255) 
);

INSERT INTO site_types (site_type_id, slug, description) VALUES (1, 'Company', 'Good for companies.');
INSERT INTO site_types (site_type_id, slug, description) VALUES (2, 'Organization', 'Good for organizations.');
INSERT INTO site_types (site_type_id, slug, description) VALUES (3, 'Workspace', 'Good for workspaces.');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS site_types;
-- +goose StatementEnd
