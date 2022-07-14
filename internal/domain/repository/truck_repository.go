package repository

import (
	"context"

	"github.com/alejandroik/trazavino-api/internal/domain/entity"
)

type TruckRepository interface {
	AddTruck(ctx context.Context, truck *entity.Truck) (*entity.Truck, error)
	GetTruck(ctx context.Context, truckId int64) (*entity.Truck, error)
}