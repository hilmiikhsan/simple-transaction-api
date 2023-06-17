package customer_address

import (
	"context"
	"errors"

	"github.com/hilmiikhsan/simple-transaction-api/entity"
	"gorm.io/gorm"
)

func NewCustomerAddressInterface(DB *gorm.DB) CustomerAddressRepositoryInterface {
	return &customerAddressRepository{
		DB: DB,
	}
}

type customerAddressRepository struct {
	*gorm.DB
}

func (customerAddressRepository *customerAddressRepository) FindByID(ctx context.Context, id, customerID int) (entity.CustomerAddress, error) {
	results := entity.CustomerAddress{}
	query := customerAddressRepository.DB.WithContext(ctx).
		Table("customer_address").
		Select("customer_address.*, customer.customer_name").
		Joins("JOIN customer ON customer.id = customer_address.customer_id").
		Where("customer_address.id = ?", id).
		Where("customer_address.customer_id = ?", customerID)

	query = query.Preload("Customer")

	result := query.Find(&results)
	if result.RowsAffected == 0 {
		return results, errors.New("Customer and address not found")
	}

	return results, nil
}
