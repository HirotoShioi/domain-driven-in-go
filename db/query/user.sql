-- name: CreateUser :one
INSERT INTO users (
    username,
    email
) VALUES (
    $1, $2
) RETURNING *;

-- name: GetUser :one
SELECT * FROM users
WHERE userid = $1
LIMIT 1;

-- name: DeleteUser :exec
DELETE FROM users
WHERE userid = $1;


-- name: UpdateUser :one
UPDATE users
SET username = $2
WHERE userid = $1
RETURNING *;
