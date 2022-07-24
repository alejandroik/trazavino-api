// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.14.0
// source: tank.sql

package generated

import (
	"context"
	"database/sql"
	"time"

	"github.com/google/uuid"
)

const addTank = `-- name: AddTank :exec
INSERT INTO tank (id, created_at, name, is_empty)
VALUES ($1, $2, $3, $4)
`

type AddTankParams struct {
	ID        uuid.UUID
	CreatedAt time.Time
	Name      string
	IsEmpty   bool
}

func (q *Queries) AddTank(ctx context.Context, arg AddTankParams) error {
	_, err := q.db.ExecContext(ctx, addTank,
		arg.ID,
		arg.CreatedAt,
		arg.Name,
		arg.IsEmpty,
	)
	return err
}

const getTank = `-- name: GetTank :one
SELECT id, created_at, updated_at, deleted_at, winery_id, name, is_empty
FROM tank
WHERE id = $1
LIMIT 1
`

func (q *Queries) GetTank(ctx context.Context, id uuid.UUID) (Tank, error) {
	row := q.db.QueryRowContext(ctx, getTank, id)
	var i Tank
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
		&i.WineryID,
		&i.Name,
		&i.IsEmpty,
	)
	return i, err
}

const listTanks = `-- name: ListTanks :many
SELECT id, created_at, updated_at, deleted_at, winery_id, name, is_empty
FROM tank
ORDER BY created_at DESC
OFFSET $1 LIMIT $2
`

type ListTanksParams struct {
	Offset int32
	Limit  int32
}

func (q *Queries) ListTanks(ctx context.Context, arg ListTanksParams) ([]Tank, error) {
	rows, err := q.db.QueryContext(ctx, listTanks, arg.Offset, arg.Limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Tank
	for rows.Next() {
		var i Tank
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.DeletedAt,
			&i.WineryID,
			&i.Name,
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

const updateTank = `-- name: UpdateTank :exec
UPDATE tank
SET name       = COALESCE($2, name),
    updated_at = COALESCE($3, updated_at),
    is_empty   = COALESCE($4, is_empty)
WHERE id = $1
`

type UpdateTankParams struct {
	ID        uuid.UUID
	Name      string
	UpdatedAt sql.NullTime
	IsEmpty   bool
}

func (q *Queries) UpdateTank(ctx context.Context, arg UpdateTankParams) error {
	_, err := q.db.ExecContext(ctx, updateTank,
		arg.ID,
		arg.Name,
		arg.UpdatedAt,
		arg.IsEmpty,
	)
	return err
}
