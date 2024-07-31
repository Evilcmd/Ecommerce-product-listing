// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: admin.sql

package postgres

import (
	"context"

	"github.com/google/uuid"
)

const addAdmin = `-- name: AddAdmin :one
INSERT INTO admin(id, username, passwd) VALUES($1, $2, $3) RETURNING id, username, passwd
`

type AddAdminParams struct {
	ID       uuid.UUID
	Username string
	Passwd   string
}

func (q *Queries) AddAdmin(ctx context.Context, arg AddAdminParams) (Admin, error) {
	row := q.db.QueryRowContext(ctx, addAdmin, arg.ID, arg.Username, arg.Passwd)
	var i Admin
	err := row.Scan(&i.ID, &i.Username, &i.Passwd)
	return i, err
}

const getAdmin = `-- name: GetAdmin :one
SELECT id, username, passwd FROM admin WHERE username=$1
`

func (q *Queries) GetAdmin(ctx context.Context, username string) (Admin, error) {
	row := q.db.QueryRowContext(ctx, getAdmin, username)
	var i Admin
	err := row.Scan(&i.ID, &i.Username, &i.Passwd)
	return i, err
}
