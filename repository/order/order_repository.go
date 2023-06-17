package order

import (
	"context"

	"github.com/hilmiikhsan/simple-transaction-api/entity"
	"gorm.io/gorm"
)

func NewOrderRepositoryInterface(DB *gorm.DB) OrderRepositoryInterface {
	return &orderRepository{
		DB: DB,
	}
}

type orderRepository struct {
	*gorm.DB
}

func (orderRepository *orderRepository) Insert(ctx context.Context, tx *gorm.DB, data entity.Order) (int, error) {
	err := tx.WithContext(ctx).Create(&data).Error
	if err != nil {
		return 0, err
	}

	return data.ID, nil
}
