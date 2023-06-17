package customer_address

import (
	"context"

	"github.com/hilmiikhsan/simple-transaction-api/entity"
)

type CustomerAddressRepositoryInterface interface {
	FindByID(ctx context.Context, id, customerID int) (entity.CustomerAddress, error)
}
