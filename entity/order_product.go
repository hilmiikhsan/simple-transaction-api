package entity

type OrderProduct struct {
	OrderID   int     `gorm:"column:order_id;type:int(11);not null"`
	ProductID int     `gorm:"column:product_id;type:int(11);not null"`
	Order     Order   `gorm:"ForeignKey:OrderID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;references:ID"`
	Product   Product `gorm:"ForeignKey:ProductID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;references:ID"`
}

func (OrderProduct) TableName() string {
	return "order_product"
}
