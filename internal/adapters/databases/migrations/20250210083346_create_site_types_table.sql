-- +goose Up
-- +goose StatementBegin
CREATE TABLE site_types (
  site_type_id SERIAL PRIMARY KEY,
  slug VARCHAR(255) NOT NULL,
  description VARCHAR(255), 
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  created_by INTEGER NOT NULL,
  updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_by INTEGER NOT NULL,
  deleted_at TIMESTAMP,

  FOREIGN KEY (created_by) REFERENCES users(user_id),
  FOREIGN KEY (updated_by) REFERENCES users(user_id)
);

INSERT INTO site_types (slug, description, created_by, updated_by) VALUES ('workspace', 'A workspace', 1, 1);
INSERT INTO site_types (slug, description, created_by, updated_by) VALUES ('company', 'A company', 1, 1);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS site_types;
-- +goose StatementEnd
