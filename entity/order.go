package entity

import "time"

type Order struct {
	ID                int       `gorm:"primaryKey;column:id;autoIncrement;type:int(11);not null"`
	CustomerID        int       `gorm:"column:customer_id;type:int(11);not null"`
	CustomerAddressID int       `gorm:"column:customer_address_id;type:int(11);not null"`
	OrderDate         time.Time `gorm:"column:order_date;type:date;not null"`
	// Products          []Product       `gorm:"many2many:order_product;"`
	// PaymentMethods    []PaymentMethod `gorm:"many2many:order_payment;"`
	Customer        Customer        `gorm:"ForeignKey:CustomerID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;references:ID"`
	CustomerAddress CustomerAddress `gorm:"ForeignKey:CustomerAddressID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;references:ID"`
}

func (Order) TableName() string {
	return "order"
}
