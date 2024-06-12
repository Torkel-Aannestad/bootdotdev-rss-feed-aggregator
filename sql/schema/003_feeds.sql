-- +goose Up
CREATE TABLE feeds (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    name VARCHAR (255) NOT NULL,
    url TEXT UNIQUE,
    user_id UUID REFERENCES users(id) ON DELETE CASCADE
);

-- +goose Down
DROP TABLE feeds;