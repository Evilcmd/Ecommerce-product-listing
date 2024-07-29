-- +goose Up
CREATE TABLE catalog(id UUID PRIMARY KEY, name TEXT NOT NULL, description TEXT NOT NULL, price INTEGER NOT NULL);

-- +goose Down
DROP TABLE catalog;