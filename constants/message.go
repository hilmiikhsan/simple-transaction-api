package constants

import "errors"

var (
	ErrCustomerAndAddressNotFound = errors.New("Customer and address not found")
	ErrProductNotFound            = errors.New("Product not found")
	ErrPaymentMethodNotFound      = errors.New("Payment method not found")
)
