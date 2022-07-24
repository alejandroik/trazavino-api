// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.14.0
// source: cask.sql

package generated

import (
	"context"
	"database/sql"
	"time"

	"github.com/google/uuid"
)

const addCask = `-- name: AddCask :exec
INSERT INTO cask (id, created_at, name, c_type, is_empty)
VALUES ($1, $2, $3, $4, $5)
`

type AddCaskParams struct {
	ID        uuid.UUID
	CreatedAt time.Time
	Name      string
	CType     string
	IsEmpty   bool
}

func (q *Queries) AddCask(ctx context.Context, arg AddCaskParams) error {
	_, err := q.db.ExecContext(ctx, addCask,
		arg.ID,
		arg.CreatedAt,
		arg.Name,
		arg.CType,
		arg.IsEmpty,
	)
	return err
}

const getCask = `-- name: GetCask :one
SELECT id, created_at, updated_at, deleted_at, name, c_type, is_empty
FROM cask
WHERE id = $1
LIMIT 1
`

func (q *Queries) GetCask(ctx context.Context, id uuid.UUID) (Cask, error) {
	row := q.db.QueryRowContext(ctx, getCask, id)
	var i Cask
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
		&i.Name,
		&i.CType,
		&i.IsEmpty,
	)
	return i, err
}

const listCasks = `-- name: ListCasks :many
SELECT id, created_at, updated_at, deleted_at, name, c_type, is_empty
FROM cask
ORDER BY created_at DESC
OFFSET $1 LIMIT $2
`

type ListCasksParams struct {
	Offset int32
	Limit  int32
}

func (q *Queries) ListCasks(ctx context.Context, arg ListCasksParams) ([]Cask, error) {
	rows, err := q.db.QueryContext(ctx, listCasks, arg.Offset, arg.Limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Cask
	for rows.Next() {
		var i Cask
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.DeletedAt,
			&i.Name,
			&i.CType,
			&i.IsEmpty,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateCask = `-- name: UpdateCask :exec
UPDATE cask
SET name       = COALESCE($2, name),
    updated_at = COALESCE($3, updated_at),
    c_type     = COALESCE($4, c_type),
    is_empty   = COALESCE($5, is_empty)
WHERE id = $1
`

type UpdateCaskParams struct {
	ID        uuid.UUID
	Name      string
	UpdatedAt sql.NullTime
	CType     string
	IsEmpty   bool
}

func (q *Queries) UpdateCask(ctx context.Context, arg UpdateCaskParams) error {
	_, err := q.db.ExecContext(ctx, updateCask,
		arg.ID,
		arg.Name,
		arg.UpdatedAt,
		arg.CType,
		arg.IsEmpty,
	)
	return err
}