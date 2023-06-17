package entity

type Customer struct {
	ID           int    `gorm:"primaryKey;column:id;autoIncrement;type:int(11);not null"`
	CustomerName string `gorm:"column:customer_name;type:varchar(255);not null"`
}

func (Customer) TableName() string {
	return "customer"
}
