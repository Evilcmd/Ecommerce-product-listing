-- +goose Up
CREATE TABLE admin(id UUID PRIMARY KEY, username TEXT UNIQUE NOT NULL, passwd TEXT NOT NULL);

-- +goose Down
DROP TABLE admin;