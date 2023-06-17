package model

type CustomerAddressModel struct {
	ID         int    `json:"id"`
	CustomerID int    `json:"customer_id"`
	Address    string `json:"address"`
}
