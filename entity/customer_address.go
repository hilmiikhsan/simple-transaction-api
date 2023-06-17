package entity

type CustomerAddress struct {
	ID         int      `gorm:"primaryKey;column:id;autoIncrement;type:int(11);not null"`
	CustomerID int      `gorm:"column:customer_id;type:int(11);not null"`
	Address    string   `gorm:"column:address;type:varchar(255);not null"`
	Customer   Customer `gorm:"ForeignKey:CustomerID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;references:ID"`
}

func (CustomerAddress) TableName() string {
	return "customer_address"
}
