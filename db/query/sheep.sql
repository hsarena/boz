
-- name: CreateSheep :one
INSERT INTO Sheeps (
  name,
  breed,
  wool,
  color
) VALUES (
  $1, $2, $3, $4
) RETURNING *;


-- name: GetSheep :one
SELECT * FROM Sheeps
WHERE id = $1 LIMIT 1;

-- name: ListSheeps :many
SELECT * FROM Sheeps
ORDER BY id
LIMIT $1
OFFSET $2;


-- name: UpdateSheep :one
UPDATE Sheeps 
SET name = $2
WHERE id = $1
RETURNING *;

-- name: DeleteSheep :exec
DELETE FROM Sheeps
WHERE id = $1;