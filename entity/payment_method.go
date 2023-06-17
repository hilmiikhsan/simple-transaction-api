package entity

type PaymentMethod struct {
	ID       int    `gorm:"primaryKey;column:id;autoIncrement;type:int(11);not null"`
	Name     string `gorm:"column:name;type:varchar(255);not null"`
	IsActive bool   `gorm:"column:is_active;type:tinyint(1);not null"`
}

func (PaymentMethod) TableName() string {
	return "payment_method"
}
