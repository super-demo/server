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

-- INSERT INTO notification_users (notification_id, user_id) VALUES (1, 1);
-- INSERT INTO notification_users (notification_id, user_id) VALUES (1, 2);
-- INSERT INTO notification_users (notification_id, user_id) VALUES (1, 3);

-- INSERT INTO notification_users (notification_id, user_id) VALUES (2, 1);
-- INSERT INTO notification_users (notification_id, user_id) VALUES (2, 2);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS notification_users;
-- +goose StatementEnd
