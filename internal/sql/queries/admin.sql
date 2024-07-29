-- name: GetAdmin :one
SELECT * FROM admin WHERE username=$1;