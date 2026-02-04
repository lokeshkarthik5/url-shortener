-- +goose Up

CREATE TABLE links (
    id UUID PRIMARY KEY,
    longUrl TEXT NOT NULL,
    short TEXT UNIQUE,
    createdAt TIMESTAMP NOT NULL,
    updatedAt TIMESTAMP NOT NULL,
    accessCount NUMBER
);

-- +goose Down
DROP TABLE IF EXISTS links;