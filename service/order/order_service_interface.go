package order

import (
	"context"

	"github.com/hilmiikhsan/simple-transaction-api/model"
)

type OrderServiceInterface interface {
	CreateOrder(ctx context.Context, order model.CreateOrderModel) error
}
