-- +goose Up
-- +goose StatementBegin
CREATE TABLE users (
  user_id SERIAL PRIMARY KEY,
  user_level_id INT,
  sub_role_id INT,
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
-- INSERT INTO users (user_level_id, name, email) VALUES (1, 'Nopnapat NORASRI', 'nopnapat.n@ku.th');
-- INSERT INTO users (user_level_id, name, email) VALUES (6, 'Nopnapat Norasri', 'nopnapatn@gmail.com');
INSERT INTO users (user_level_id, name, email) VALUES (1, 'Thanyamas Chancharoen', 'thanyamas.c@ku.th');
INSERT INTO users (user_level_id, name, email) VALUES
(3, 'Sutthiphong Thiangtham', 'sutthiphong.t@ku.th'),
(3, 'Anuchit Chaichan', 'anuchit.c@ku.th'),
(4, 'Kanokwan Srisuk', 'kanokwan.s@ku.th'),
(4, 'Phannee Wongjira', 'phannee.w@ku.th'),
(4, 'Chatchawal Santisuk', 'chatchawal.s@ku.th'),
(4, 'Sirinapa Nakathum', 'sirinapa.n@ku.th'),
(4, 'Teerawat Methee', 'teerawat.m@ku.th'),
(5, 'Phanuwat Rattanasuk', 'phanuwat.r@ku.th'),
(5, 'Mintra Sreeboon', 'mintra.s@ku.th'),
(6, 'Uthai Phrommin', 'uthai.p@ku.th'),
(6, 'Wararat Saengtham', 'wararat.s@ku.th'),
(6, 'Chaiyapol Nakachat', 'chaiyapol.n@ku.th'),
(6, 'Supaporn Amornwit', 'supaporn.a@ku.th'),
(6, 'Worachai Chanthap', 'worachai.c@ku.th'),
(6, 'Prapasorn Pathanawit', 'prapasorn.p@ku.th'),
(6, 'Duangjai Phromchat', 'duangjai.p@ku.th'),
(6, 'Laddawan Pimprasith', 'laddawan.p@ku.th'),
(6, 'Chanakan Rattanajaroen', 'chanakan.r@ku.th'),
(6, 'Patcharapa Chuenjai', 'patcharapa.c@ku.th'),
(6, 'Pimnipa Watthanatham', 'pimnipa.w@ku.th');

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS users;
-- +goose StatementEnd
