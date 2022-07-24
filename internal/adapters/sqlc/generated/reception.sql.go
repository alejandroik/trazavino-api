// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.14.0
// source: reception.sql

package generated

import (
	"context"
	"time"

	"github.com/google/uuid"
)

const addReception = `-- name: AddReception :exec
INSERT INTO reception (id, created_at, weight, sugar, truck_id, vineyard_id, grape_type_id)
VALUES ($1, $2, $3, $4, $5, $6, $7)
`

type AddReceptionParams struct {
	ID          uuid.UUID
	CreatedAt   time.Time
	Weight      int32
	Sugar       int32
	TruckID     uuid.UUID
	VineyardID  uuid.UUID
	GrapeTypeID uuid.UUID
}

func (q *Queries) AddReception(ctx context.Context, arg AddReceptionParams) error {
	_, err := q.db.ExecContext(ctx, addReception,
		arg.ID,
		arg.CreatedAt,
		arg.Weight,
		arg.Sugar,
		arg.TruckID,
		arg.VineyardID,
		arg.GrapeTypeID,
	)
	return err
}

const getReception = `-- name: GetReception :one
SELECT id, created_at, updated_at, deleted_at, weight, sugar, truck_id, vineyard_id, grape_type_id
FROM reception
WHERE id = $1
LIMIT 1
`

func (q *Queries) GetReception(ctx context.Context, id uuid.UUID) (Reception, error) {
	row := q.db.QueryRowContext(ctx, getReception, id)
	var i Reception
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
		&i.Weight,
		&i.Sugar,
		&i.TruckID,
		&i.VineyardID,
		&i.GrapeTypeID,
	)
	return i, err
}

const listReceptions = `-- name: ListReceptions :many
SELECT id, created_at, updated_at, deleted_at, weight, sugar, truck_id, vineyard_id, grape_type_id
FROM reception
ORDER BY created_at DESC
OFFSET $1 LIMIT $2
`

type ListReceptionsParams struct {
	Offset int32
	Limit  int32
}

func (q *Queries) ListReceptions(ctx context.Context, arg ListReceptionsParams) ([]Reception, error) {
	rows, err := q.db.QueryContext(ctx, listReceptions, arg.Offset, arg.Limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Reception
	for rows.Next() {
		var i Reception
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.DeletedAt,
			&i.Weight,
			&i.Sugar,
			&i.TruckID,
			&i.VineyardID,
			&i.GrapeTypeID,
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
