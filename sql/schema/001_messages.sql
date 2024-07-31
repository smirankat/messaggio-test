-- +goose Up

CREATE TABLE msgs (
    id UUID PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    message TEXT NOT NULL
);

-- +goose Down
DROP TABLE msgs;
