-- +goose Up
-- +goose StatementBegin
CREATE TABLE site_trees (
    site_tree_id SERIAL PRIMARY KEY,
    site_parent_id INTEGER NOT NULL,
    site_child_id INTEGER NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    created_by INTEGER NOT NULL,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_by INTEGER NOT NULL,
    deleted_at TIMESTAMP,

    FOREIGN KEY (site_parent_id) REFERENCES sites(site_id),
    FOREIGN KEY (site_child_id) REFERENCES sites(site_id),
    FOREIGN KEY (created_by) REFERENCES users(user_id),
    FOREIGN KEY (updated_by) REFERENCES users(user_id)
);

INSERT INTO site_trees (site_parent_id, site_child_id, created_by, updated_by) VALUES (1, 2, 1, 1);
INSERT INTO site_trees (site_parent_id, site_child_id, created_by, updated_by) VALUES (1, 3, 1, 1);
INSERT INTO site_trees (site_parent_id, site_child_id, created_by, updated_by) VALUES (1, 4, 1, 1);
INSERT INTO site_trees (site_parent_id, site_child_id, created_by, updated_by) VALUES (2, 5, 1, 1);
INSERT INTO site_trees (site_parent_id, site_child_id, created_by, updated_by) VALUES (2, 6, 1, 1);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS site_trees;
-- +goose StatementEnd
