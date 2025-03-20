-- +goose Up
-- +goose StatementBegin
CREATE TABLE notifications (
  notification_id SERIAL PRIMARY KEY,
  site_id INTEGER NOT NULL,
  action VARCHAR(255) NOT NULL,
  detail TEXT,
  image_url TEXT,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  created_by INTEGER NOT NULL,

  FOREIGN KEY (site_id) REFERENCES sites(site_id),
  FOREIGN KEY (created_by) REFERENCES users(user_id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS notifications;
-- +goose StatementEnd
