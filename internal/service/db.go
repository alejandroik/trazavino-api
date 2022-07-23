package service

import (
	"context"
	"os"

	"github.com/alejandroik/trazavino/internal/adapters/dynamodb"
	"github.com/alejandroik/trazavino/internal/adapters/sqlc"
	"github.com/alejandroik/trazavino/internal/domain/repository"
)

type repositories struct {
	ReceptionRepository    repository.ReceptionRepository
	MacerationRepository   repository.MacerationRepository
	FermentationRepository repository.FermentationRepository
}

func getRepositories(ctx context.Context) (*repositories, error) {
	r := &repositories{}

	dbAdapter := os.Getenv("DB_ADAPTER")
	if dbAdapter == "" {
		dbAdapter = "postgres"
	}

	switch dbAdapter {
	case "postgres":
		db, err := sqlc.NewPostgresConnection()
		if err != nil {
			return nil, err
		}

		r.ReceptionRepository = sqlc.NewReceptionRepository(db)
		r.MacerationRepository = sqlc.NewMacerationRepository(db)
		r.FermentationRepository = sqlc.NewFermentationRepository(db)

	case "dynamodb":
		db, err := dynamodb.NewDynamoDbClient(ctx)
		if err != nil {
			return nil, err
		}

		r.ReceptionRepository = dynamodb.NewReceptionDynamoDbRepository(db)
		r.MacerationRepository = dynamodb.NewMacerationDynamodbRepository(db)
	}

	return r, nil
}
