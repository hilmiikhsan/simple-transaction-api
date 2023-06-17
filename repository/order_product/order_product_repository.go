package order_product

import (
	"context"

	"github.com/hilmiikhsan/simple-transaction-api/entity"
	"gorm.io/gorm"
)

func NewOrderProductRepositoryInterface(DB *gorm.DB) OrderProductRepositoryInterface {
	return &orderProductRepository{
		DB: DB,
	}
}

type orderProductRepository struct {
	*gorm.DB
}

func (orderProductRepository *orderProductRepository) Insert(ctx context.Context, tx *gorm.DB, data entity.OrderProduct) error {
	err := tx.WithContext(ctx).Create(&data).Error
	if err != nil {
		return err
	}

	return nil
}

func (orderProductRepository *orderProductRepository) InsertBatch(ctx context.Context, tx *gorm.DB, data []entity.OrderProduct) error {
	err := tx.WithContext(ctx).Create(&data).Error
	if err != nil {
		return err
	}

	return nil
}
