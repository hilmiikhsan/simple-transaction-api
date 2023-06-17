package model

type CreateOrderModel struct {
	CustomerID        int                 `json:"customer_id"`
	CustomerAddressID int                 `json:"customer_address_id"`
	OrderProducts     []OrderProductModel `json:"order_products"`
	OrderPayments     []OrderPaymentModel `json:"order_payments"`
}
