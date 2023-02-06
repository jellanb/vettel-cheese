package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type Order struct {
	OrderNumber  int
	Date         primitive.DateTime
	SellerName   string
	Status       string
	SaleAmount   int
	Products     []Product
	Dispatch     Dispatch
	Payment      Payment
	ReasonSocial string
}

type Dispatch struct {
	CreatedDate       primitive.DateTime
	CreditTime        primitive.DateTime
	ReasonSocial      string
	Address           string
	Commune           string
	AddressNumber     int
	Region            string
	AdministratorName string
	Phone             string
	DeliveryDate      primitive.DateTime
	Status            string
	DeliveryName      string
	Amount            int
	Products          []Product
}

type Payment struct {
	Id                int
	PaymentDate       primitive.DateTime
	CreditTime        primitive.DateTime
	Amount            int
	Status            string
	Folio             string
	SummaryPaidAmount int
	ReasonSocial      string
	NetAmount         int
	Iva               int
	DocumentDate      primitive.DateTime
	CreditPaid        []Quota
}

type Quota struct {
	Amount       int
	Different    int
	Date         primitive.DateTime
	ReceptorName string
}
