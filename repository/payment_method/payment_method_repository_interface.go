package payment_method

import (
	"context"

	"github.com/hilmiikhsan/simple-transaction-api/entity"
)

type PaymentMenthodRepositoryInterface interface {
	FindBatchByID(ctx context.Context, paymentMethodIDs []int) ([]entity.PaymentMethod, error)
}
