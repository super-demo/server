-- +goose Up
-- +goose StatementBegin
CREATE TABLE announcements (
  announcement_id SERIAL PRIMARY KEY,
  site_id INTEGER NOT NULL,
  title VARCHAR(255) NOT NULL,
  short_description TEXT,
  image_url TEXT,
  link_url TEXT,
  is_pin BOOLEAN NOT NULL DEFAULT FALSE,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  created_by INTEGER NOT NULL,
  deleted_at TIMESTAMP,

  FOREIGN KEY (site_id) REFERENCES sites(site_id),
  FOREIGN KEY (created_by) REFERENCES users(user_id)
);

INSERT INTO announcements (site_id, title, short_description, image_url, link_url, is_pin, created_by) VALUES (1, 'Welcome to the new site!', 'We are excited to announce the launch of our new site!', 'https://scontent-bkk1-2.xx.fbcdn.net/v/t39.30808-6/476509205_555409170843779_2189911961991409505_n.jpg?stp=dst-jpg_s1080x2048_tt6&_nc_cat=107&ccb=1-7&_nc_sid=cc71e4&_nc_ohc=dvVyY7fKDE4Q7kNvgEUsxY5&_nc_oc=Adm7pMZkVfAy3GQbposPMK9VpKk81CyRhNdrr2JqPgAp0TnQlYyuOJT7nelQ2oDVmBM&_nc_zt=23&_nc_ht=scontent-bkk1-2.xx&_nc_gid=eBOCDkseztFEgrAKMda5Xg&oh=00_AYGH3QWile1EBMI5YsbJZHdZDO8Ut6TAyNEjyCzqITjDEw&oe=67DF7CDD', 'https://nopnapatn.dev', TRUE, 1);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS announcements;
-- +goose StatementEnd
