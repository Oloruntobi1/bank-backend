-- name: CreateClient :one
INSERT INTO clients (
  name, email, password
) VALUES (
  $1, $2, $3
)
RETURNING *;

-- name: GetClient :one
SELECT * FROM clients
WHERE id = $1 LIMIT 1;

-- name: ListClients :many
SELECT * FROM clients
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: DeleteClient :exec
DELETE FROM clients WHERE id = $1;