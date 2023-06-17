package entity

type OrderPayment struct {
	OrderID         int           `gorm:"column:order_id;type:int(11);not null"`
	PaymentMethodID int           `gorm:"column:payment_method_id;type:int(11);not null"`
	Order           Order         `gorm:"ForeignKey:OrderID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;references:ID"`
	PaymentMethod   PaymentMethod `gorm:"ForeignKey:PaymentMethodID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;references:ID"`
}

func (OrderPayment) TableName() string {
	return "order_payment"
}
