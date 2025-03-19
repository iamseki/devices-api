-- name: GetDevice :one
SELECT id, name, brand, state, creation_time FROM devices WHERE id = $1;

-- name: ListDevices :many
SELECT 
  id, name, brand, state, creation_time
FROM 
  devices
WHERE
  (brand = $1 OR $1 = '') AND
  (state = $2 OR $2 = '') AND
  (name = $3 OR $3 = '');

-- name: InsertDevice :exec
INSERT INTO devices (name, brand) VALUES ($1, $2);

-- name: DeleteDevice :exec
DELETE FROM devices WHERE id = $1;

-- name: UpdateDevice :exec
UPDATE devices SET name = $2, brand = $3, state = $4 WHERE id = $1;