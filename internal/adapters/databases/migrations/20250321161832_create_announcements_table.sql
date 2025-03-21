-- +goose Up
-- +goose StatementBegin
CREATE TABLE announcements (
  announcement_id SERIAL PRIMARY KEY,
  site_id INTEGER NOT NULL,
  title VARCHAR(255) NOT NULL,
  short_description TEXT,
  image_url TEXT,
  is_pin BOOLEAN NOT NULL DEFAULT FALSE,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  created_by INTEGER NOT NULL,

  FOREIGN KEY (site_id) REFERENCES sites(site_id),
  FOREIGN KEY (created_by) REFERENCES users(user_id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS announcements;
-- +goose StatementEnd
