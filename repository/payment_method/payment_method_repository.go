package payment_method

import (
	"context"
	"errors"

	"github.com/hilmiikhsan/simple-transaction-api/entity"
	"gorm.io/gorm"
)

func NewPaymentMethodRepositoryInterface(DB *gorm.DB) PaymentMenthodRepositoryInterface {
	return &paymentMethodRepository{
		DB: DB,
	}
}

type paymentMethodRepository struct {
	*gorm.DB
}

func (paymentMethodRepository *paymentMethodRepository) FindBatchByID(ctx context.Context, paymentMethodIDs []int) ([]entity.PaymentMethod, error) {
	results := []entity.PaymentMethod{}
	err := paymentMethodRepository.DB.Where("id IN ?", paymentMethodIDs).WithContext(ctx).Find(&results).Error
	if err != nil {
		return results, err
	}

	if len(results) == 0 {
		return results, errors.New("Payment Method not found")
	}

	return results, nil
}
