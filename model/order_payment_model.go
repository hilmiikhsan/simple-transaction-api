package model

type OrderPaymentModel struct {
	OrderID         int `json:"order_id"`
	PaymentMethodID int `json:"payment_method_id"`
}
