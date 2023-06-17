package order_product

import (
	"context"

	"github.com/hilmiikhsan/simple-transaction-api/entity"
	"gorm.io/gorm"
)

type OrderProductRepositoryInterface interface {
	InsertBatch(ctx context.Context, tx *gorm.DB, data []entity.OrderProduct) error
}
