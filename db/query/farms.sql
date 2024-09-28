-- name: CreateFarm :one
INSERT INTO farms (
  farmer_id,
  address,
  size,
  government_id
) VALUES (
  $1, $2, $3, $4
)
RETURNING *;

-- name: GetFarm :one
SELECT * FROM farms
WHERE farmer_id = $1;

-- name: ListFarms :many
SELECT * FROM farms
WHERE farmer_id = $1
ORDER BY id
LIMIT $2
OFFSET $3;

-- name: UpdateFarm :one
UPDATE farms SET
    address = COALESCE(NULLIF($2, ''), address),
    size = COALESCE(NULLIF($3, ''), size),
    government_id = COALESCE(NULLIF($4, ''), government_id)
WHERE id = $1
RETURNING *;

-- name: DeleteFarm :exec
DELETE FROM farms
WHERE id = $1;