package order

import (
	"context"

	"github.com/hilmiikhsan/simple-transaction-api/entity"
	"gorm.io/gorm"
)

type OrderRepositoryInterface interface {
	Insert(ctx context.Context, tx *gorm.DB, data entity.Order) (int, error)
}
