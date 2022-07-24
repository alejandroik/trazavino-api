package sqlc

import (
	"context"
	"database/sql"
	"time"

	"github.com/alejandroik/trazavino/internal/adapters/sqlc/generated"
	"github.com/alejandroik/trazavino/internal/domain/entity"
	"github.com/alejandroik/trazavino/internal/domain/entity/enum/process_type"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type AgeingRepository struct {
	db *sqlx.DB
}

func NewAgeingRepository(db *sqlx.DB) *AgeingRepository {
	if db == nil {
		panic("missing db")
	}

	return &AgeingRepository{db: db}
}

func (r AgeingRepository) AddAgeing(ctx context.Context, a *entity.Ageing) error {
	ageUuid, err := uuid.Parse(a.UUID())
	if err != nil {
		return nil
	}
	tankUuid, err := uuid.Parse(a.TankUUID())
	if err != nil {
		return err
	}
	caskUuid, err := uuid.Parse(a.CaskUUID())
	if err != nil {
		return err
	}

	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	q := generated.New(tx)

	now := time.Now()

	fer, err := q.FindFermentation(ctx, tankUuid)
	if err != nil {
		tx.Rollback()
		return err
	}

	if err = q.UpdateProcess(ctx, generated.UpdateProcessParams{
		ID:        fer.ID,
		UpdatedAt: sql.NullTime{Time: now, Valid: true},
		EndTime:   sql.NullTime{Time: a.StartTime(), Valid: true},
	}); err != nil {
		tx.Rollback()
		return err
	}

	if err = q.UpdateTank(ctx, generated.UpdateTankParams{
		ID:        tankUuid,
		UpdatedAt: sql.NullTime{Time: now, Valid: true},
		IsEmpty:   true,
	}); err != nil {
		tx.Rollback()
		return err
	}

	if err = q.UpdateCask(ctx, generated.UpdateCaskParams{
		ID:        caskUuid,
		UpdatedAt: sql.NullTime{Time: now, Valid: true},
		IsEmpty:   false,
	}); err != nil {
		tx.Rollback()
		return err
	}

	if err = q.AddProcess(ctx, generated.AddProcessParams{
		ID:         ageUuid,
		CreatedAt:  now,
		StartTime:  a.StartTime(),
		PType:      process_type.Ageing.String(),
		PreviousID: uuid.NullUUID{UUID: fer.ID, Valid: true},
	}); err != nil {
		tx.Rollback()
		return err
	}

	if err = q.AddAgeing(ctx, generated.AddAgeingParams{
		ID:        ageUuid,
		CreatedAt: now,
		TankID:    tankUuid,
		CaskID:    caskUuid,
	}); err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}

func (r AgeingRepository) GetAgeing(ctx context.Context, ageingUUID string) (*entity.Ageing, error) {
	return nil, nil
}

func (r AgeingRepository) UpdateAgeing(
	ctx context.Context,
	ageingUUID string,
	updateFn func(ctx context.Context, f *entity.Ageing) (*entity.Ageing, error),
) error {
	return nil
}