-- name: GetDevice :one
SELECT id, name, brand, state, created_at FROM devices WHERE id = $1;

-- name: ListDevices :many
SELECT 
  id, name, brand, state, created_at 
FROM 
  devices
WHERE
  brand = $1 OR state = $2;

-- name: InsertDevice :exec
INSERT INTO devices (name, brand, state) VALUES ($1, $2, $3);

-- name: DeleteDevice :exec
DELETE FROM devices WHERE id = $1;

-- name: UpdateDevice :exec
UPDATE devices SET name = $2, brand = $3, state = $4 WHERE id = $1;