
-- name: CreateEntry :one
INSERT INTO enteries (
    account_id,
    amount 
) VALUES (
    $1, $2
)
 RETURNING * ;

-- name: GetEntry :one
SELECT * FROM enteries
WHERE id = $1 LIMIT 1 ;

-- name: ListEntry :many
SELECT * FROM enteries
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: UpdateEntry :one
UPDATE enteries
SET amount = $2
WHERE id = $1
RETURNING *;

-- name: DeleteEntry :exec
DELETE FROM enteries
WHERE id = $1;
