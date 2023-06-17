package product

import (
	"context"
	"errors"

	"github.com/hilmiikhsan/simple-transaction-api/entity"
	"gorm.io/gorm"
)

func NewProductRepositoryInterface(DB *gorm.DB) ProductRepositoryInterface {
	return &productRepository{
		DB: DB,
	}
}

type productRepository struct {
	*gorm.DB
}

func (productRepository *productRepository) FindBatchByID(ctx context.Context, productIDs []int) ([]entity.Product, error) {
	results := []entity.Product{}
	err := productRepository.DB.Where("id IN ?", productIDs).WithContext(ctx).Find(&results).Error
	if err != nil {
		return results, err
	}

	if len(results) == 0 {
		return results, errors.New("Product not found")
	}

	return results, nil
}
