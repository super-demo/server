-- +goose Up
-- +goose StatementBegin
CREATE TABLE site_logs (
  site_log_id SERIAL PRIMARY KEY,
  site_id INTEGER NOT NULL,
  action VARCHAR(255) NOT NULL,
  detail TEXT,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  created_by INTEGER NOT NULL,

  FOREIGN KEY (site_id) REFERENCES sites(site_id),
  FOREIGN KEY (created_by) REFERENCES users(user_id)
);

INSERT INTO site_logs (site_id, action, detail, created_by) VALUES (1, 'create', 'Create a site', 1);
INSERT INTO site_logs (site_id, action, detail, created_by) VALUES (1, 'update', 'Update a site', 1);

INSERT INTO site_logs (site_id, action, detail, created_by) VALUES (2, 'create', 'Create a site', 1);
INSERT INTO site_logs (site_id, action, detail, created_by) VALUES (2, 'update', 'Update a site', 1);

INSERT INTO site_logs (site_id, action, detail, created_by) VALUES (3, 'create', 'Create a site', 1);

INSERT INTO site_logs (site_id, action, detail, created_by) VALUES (4, 'create', 'Create a site', 1);

INSERT INTO site_logs (site_id, action, detail, created_by) VALUES (5, 'create', 'Create a site', 1);

INSERT INTO site_logs (site_id, action, detail, created_by) VALUES (6, 'create', 'Create a site', 1);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS site_logs;
-- +goose StatementEnd
