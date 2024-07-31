-- name: GetAllProducts :many
SELECT * FROM catalog;

-- name: GetProduct :one
SELECT * FROM catalog WHERE id=$1;

-- name: AddProduct :one
INSERT INTO catalog (id, name, description, price)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: UpdateProductName :exec
UPDATE catalog SET name=$2 WHERE id=$1;

-- name: UpdateProductDescription :exec
UPDATE catalog SET description=$2 WHERE id=$1;

-- name: UpdateProductPrice :exec
UPDATE catalog SET price=$2 WHERE id=$1;

-- name: DeleteProduct :exec
DELETE FROM catalog WHERE id=$1;