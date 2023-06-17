package order_payment

import (
	"context"

	"github.com/hilmiikhsan/simple-transaction-api/entity"
	"gorm.io/gorm"
)

func NewOrderPaymentRepositoryInterface(DB *gorm.DB) OrderPaymentRepositoryInterface {
	return &orderPaymentRepository{
		DB: DB,
	}
}

type orderPaymentRepository struct {
	*gorm.DB
}

func (orderPaymentRepository *orderPaymentRepository) InsertBatch(ctx context.Context, tx *gorm.DB, data []entity.OrderPayment) error {
	err := tx.WithContext(ctx).Create(&data).Error
	if err != nil {
		return err
	}

	return nil
}
