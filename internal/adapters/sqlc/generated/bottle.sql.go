// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.14.0
// source: bottle.sql

package generated

import (
	"context"
	"database/sql"
	"time"

	"github.com/google/uuid"
)

const addBottle = `-- name: AddBottle :exec
INSERT INTO bottle (id, created_at, winery_id, name)
VALUES ($1, $2, $3, $4)
`

type AddBottleParams struct {
	ID        uuid.UUID `db:"id"`
	CreatedAt time.Time `db:"created_at"`
	WineryID  uuid.UUID `db:"winery_id"`
	Name      string    `db:"name"`
}

func (q *Queries) AddBottle(ctx context.Context, arg AddBottleParams) error {
	_, err := q.db.Exec(ctx, addBottle,
		arg.ID,
		arg.CreatedAt,
		arg.WineryID,
		arg.Name,
	)
	return err
}

const getBottle = `-- name: GetBottle :one
SELECT id, created_at, updated_at, deleted_at, winery_id, name
FROM bottle
WHERE id = $1
LIMIT 1
`

func (q *Queries) GetBottle(ctx context.Context, id uuid.UUID) (Bottle, error) {
	row := q.db.QueryRow(ctx, getBottle, id)
	var i Bottle
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
		&i.WineryID,
		&i.Name,
	)
	return i, err
}

const listBottles = `-- name: ListBottles :many
SELECT id, created_at, updated_at, deleted_at, winery_id, name
FROM bottle
ORDER BY created_at DESC
    OFFSET $1 LIMIT $2
`

type ListBottlesParams struct {
	Offset int32 `db:"offset"`
	Limit  int32 `db:"limit"`
}

func (q *Queries) ListBottles(ctx context.Context, arg ListBottlesParams) ([]Bottle, error) {
	rows, err := q.db.Query(ctx, listBottles, arg.Offset, arg.Limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Bottle
	for rows.Next() {
		var i Bottle
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.DeletedAt,
			&i.WineryID,
			&i.Name,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateBottle = `-- name: UpdateBottle :exec
UPDATE bottle
SET name       = COALESCE($2, name),
    updated_at = COALESCE($3, updated_at)
WHERE id = $1
`

type UpdateBottleParams struct {
	ID        uuid.UUID    `db:"id"`
	Name      string       `db:"name"`
	UpdatedAt sql.NullTime `db:"updated_at"`
}

func (q *Queries) UpdateBottle(ctx context.Context, arg UpdateBottleParams) error {
	_, err := q.db.Exec(ctx, updateBottle, arg.ID, arg.Name, arg.UpdatedAt)
	return err
}
