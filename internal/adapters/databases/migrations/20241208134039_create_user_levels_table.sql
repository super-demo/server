-- +goose Up
-- +goose StatementBegin
CREATE TABLE user_levels (
  user_level_id SERIAL PRIMARY KEY,
  slug VARCHAR(255) NOT NULL,
  description VARCHAR(255) 
);

INSERT INTO user_levels (user_level_id, slug, description) VALUES (1, 'Root', 'Good for people who manage everything.');
INSERT INTO user_levels (user_level_id, slug, description) VALUES (2, 'Staff', 'Good for people who manage something.');

INSERT INTO user_levels (user_level_id, slug, description) VALUES (3, 'Owner ', 'Good for people who own everything.');
INSERT INTO user_levels (user_level_id, slug, description) VALUES (4, 'Super Admin', 'Good for people who can manage everything.');
INSERT INTO user_levels (user_level_id, slug, description) VALUES (5, 'Admin', 'Good for people who just need to manage something.');
INSERT INTO user_levels (user_level_id, slug, description) VALUES (6, 'Member', 'Good for people who just need to view something.');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS user_levels;
-- +goose StatementEnd
