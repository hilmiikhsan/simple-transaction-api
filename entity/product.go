package entity

type Product struct {
	ID    int     `gorm:"primaryKey;column:id;autoIncrement;type:int(11);not null"`
	Name  string  `gorm:"column:name;type:varchar(255);not null"`
	Price float64 `gorm:"column:price;type:decimal(10,2);not null"`
}

func (Product) TableName() string {
	return "product"
}
