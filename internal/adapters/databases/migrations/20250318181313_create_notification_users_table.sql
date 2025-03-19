-- +goose Up
-- +goose StatementBegin
CREATE TABLE notification_users (
  notification_user_id SERIAL PRIMARY KEY,
  notification_id INTEGER NOT NULL,
  user_id INTEGER NOT NULL,
  is_read BOOLEAN NOT NULL DEFAULT FALSE,

  FOREIGN KEY (notification_id) REFERENCES notifications(notification_id),
  FOREIGN KEY (user_id) REFERENCES users(user_id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS notification_users;
-- +goose StatementEnd
