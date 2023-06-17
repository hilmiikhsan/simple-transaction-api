package order_payment

import (
	"context"

	"github.com/hilmiikhsan/simple-transaction-api/entity"
	"gorm.io/gorm"
)

type OrderPaymentRepositoryInterface interface {
	InsertBatch(ctx context.Context, tx *gorm.DB, data []entity.OrderPayment) error
}
