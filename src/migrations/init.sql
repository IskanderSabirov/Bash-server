
CREATE TABLE IF NOT EXISTS executed_scripts
(
    script TEXT PRIMARY KEY,
    result TEXT

);

CREATE TABLE IF NOT EXISTS users
(
    login    TEXT PRIMARY KEY,
    password TEXT
);

INSERT INTO users
VALUES ('user', 'user')
ON CONFLICT DO NOTHING;
