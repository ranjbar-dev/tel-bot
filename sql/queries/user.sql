-- name: AllUsers :many
SELECT * FROM users;

-- name: FindUser :one
SELECT * FROM users WHERE chat_id = ?;

-- name: InsertUser :one
INSERT INTO users (chat_id, name, created_at) VALUES (?, ?, ?)
RETURNING *;

-- name: UpdateUserInformation :one
UPDATE users SET name = ? WHERE chat_id = ? 
RETURNING *;

-- name: DeleteUser :exec
DELETE FROM users WHERE chat_id = ?;

