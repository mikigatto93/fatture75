package model

type ServiceExpense struct {
	Description string
	Type        string
	Price       float32
	VatCode     vatCode
}
