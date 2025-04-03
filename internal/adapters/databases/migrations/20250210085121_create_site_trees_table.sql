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

INSERT INTO site_trees (site_parent_id, site_child_id, created_by, updated_by) VALUES 
    (1, 2, 1, 1),  -- Kasetsart University -> Bangkhen Campus
    (1, 3, 1, 1),  -- Kasetsart University -> Kamphaeng Saen Campus
    (1, 4, 1, 1),  -- Kasetsart University -> Sriracha Campus
    (2, 5, 1, 1),  -- Bangkhen Campus -> Science Faculty
    (2, 6, 1, 1),  -- Bangkhen Campus -> Engineering Faculty
    (2, 7, 1, 1),  -- Bangkhen Campus -> Agriculture Faculty
    (2, 8, 1, 1),  -- Bangkhen Campus -> Economics Faculty
    (2, 9, 1, 1),  -- Bangkhen Campus -> Fisheries Faculty
    (2, 10, 1, 1), -- Bangkhen Campus -> Forestry Faculty
    (5, 11, 1, 1), -- Science Faculty -> Computer Science Department
    (5, 12, 1, 1), -- Science Faculty -> Physics Lab
    (5, 13, 1, 1), -- Science Faculty -> Chemistry Lab
    (5, 14, 1, 1), -- Science Faculty -> Biology Research Center
    (11, 15, 1, 1), -- Computer Science Department -> CS Office
    (11, 16, 1, 1); -- Computer Science Department -> CS Nisit
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS site_trees;
-- +goose StatementEnd
