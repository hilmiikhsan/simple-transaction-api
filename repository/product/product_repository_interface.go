package product

import (
	"context"

	"github.com/hilmiikhsan/simple-transaction-api/entity"
)

type ProductRepositoryInterface interface {
	FindBatchByID(ctx context.Context, productIDs []int) ([]entity.Product, error)
}
