-- name: CreateMessage :one
INSERT INTO msgs (id, created_at, updated_at, message)
VALUES (
    $1,
    $2,
    $3,
    $4
    )
RETURNING *;
