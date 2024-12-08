-- +goose Up
-- +goose StatementBegin
CREATE TABLE users (
  user_id SERIAL PRIMARY KEY,
  google_token TEXT,
  avatar_url TEXT,
  name VARCHAR(255) NOT NULL,
  nickname VARCHAR(255),
  email VARCHAR(255) UNIQUE NOT NULL,
  phone_number VARCHAR(255),
  roles VARCHAR(255),
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS users;
-- +goose StatementEnd
