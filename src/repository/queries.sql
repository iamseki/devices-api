-- name: GetDevice :one
SELECT * FROM devices WHERE id = $1;