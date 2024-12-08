-- +goose Up
-- +goose StatementBegin
CREATE TABLE user_levels (
  user_level_id SERIAL PRIMARY KEY,
  slug VARCHAR(255) NOT NULL,
  description TEXT,
);

INSERT INTO user_levels (user_level_id, user_level_title, description) VALUES (1, 'Member', 'Good for people who just need to redeem reward.');
INSERT INTO user_levels (user_level_id, user_level_title, description) VALUES (2, 'Admin', 'Good for people who just need to manage something.');
INSERT INTO user_levels (user_level_id, user_level_title, description) VALUES (3, 'Super Admin', 'Good for people who can manage everything.');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS user_levels;
-- +goose StatementEnd
