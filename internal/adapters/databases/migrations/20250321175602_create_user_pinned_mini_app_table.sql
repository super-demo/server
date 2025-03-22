-- +goose Up
-- +goose StatementBegin
CREATE TABLE user_pinned_mini_app (
    user_pinned_mini_app_id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL,
    site_mini_app_id INTEGER NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,

    FOREIGN KEY (user_id) REFERENCES users(user_id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS user_pinned_mini_app;
-- +goose StatementEnd
