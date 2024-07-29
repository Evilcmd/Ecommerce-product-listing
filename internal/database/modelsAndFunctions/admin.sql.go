// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: admin.sql

package modelsAndFunctions

import (
	"context"
)

const getAdmin = `-- name: GetAdmin :one
SELECT id, username, passwd FROM admin WHERE username=$1
`

func (q *Queries) GetAdmin(ctx context.Context, username string) (Admin, error) {
	row := q.db.QueryRowContext(ctx, getAdmin, username)
	var i Admin
	err := row.Scan(&i.ID, &i.Username, &i.Passwd)
	return i, err
}