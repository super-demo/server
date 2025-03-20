-- +goose Up
-- +goose StatementBegin
CREATE TABLE users (
  user_id SERIAL PRIMARY KEY,
  user_level_id INT,
  site_id INT,
  google_token TEXT,
  avatar_url TEXT,
  name VARCHAR(255) NOT NULL,
  nickname VARCHAR(255),
  email VARCHAR(255) UNIQUE NOT NULL,
  phone_number VARCHAR(255),
  birth_date DATE,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,

  FOREIGN KEY (user_level_id) REFERENCES user_levels(user_level_id)
);

INSERT INTO users (user_level_id, name, email) VALUES (1, 'Root', 'root@localhost');
INSERT INTO users (user_level_id, name, email) VALUES (1, 'Nopnapat Norasri', 'nopnapatn@gmail.com');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS users;
-- +goose StatementEnd
