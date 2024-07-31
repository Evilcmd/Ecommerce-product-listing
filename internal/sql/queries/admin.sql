-- name: GetAdmin :one
SELECT * FROM admin WHERE username=$1;

-- name: AddAdmin :one
INSERT INTO admin(id, username, passwd) VALUES($1, $2, $3) RETURNING *;