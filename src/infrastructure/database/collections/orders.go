package collections

import "go.mongodb.org/mongo-driver/bson/primitive"

type Order struct {
	OrderNumber  int                `bson:"orderNumber"`
	Date         primitive.DateTime `bson:"date"`
	SellerName   string             `bson:"sellerName"`
	Status       string             `bson:"status"`
	SaleAmount   int                `bson:"saleAmount"`
	Products     []Product          `bson:"products"`
	Dispatch     Dispatch           `bson:"dispatch"`
	Payment      Payment            `bson:"payment"`
	ReasonSocial string             `bson:"reasonSocial"`
}

type Dispatch struct {
	Id                int                `bson:"id"`
	CreatedDate       primitive.DateTime `bson:"createdDate"`
	ReasonSocial      string             `bson:"reasonSocial"`
	Address           string             `bson:"address"`
	Commune           string             `bson:"commune"`
	AddressNumber     int                `bson:"addressNumber"`
	Region            string             `bson:"region"`
	AdministratorName string             `bson:"administratorName"`
	Phone             string             `bson:"phone"`
	DeliveryDate      primitive.DateTime `bson:"deliveryDate"`
	Status            string             `bson:"status"`
	DeliveryName      string             `bson:"deliveryName"`
	Products          []Product          `bson:"products"`
	CreditTime        primitive.DateTime `bson:"creditTime"`
}

type Payment struct {
	Id                int                `bson:"id"`
	PaymentDate       primitive.DateTime `bson:"paymentDate"`
	CreditTime        primitive.DateTime `bson:"creditTime"`
	Amount            int                `bson:"amount"`
	Status            string             `bson:"status"`
	Folio             string             `bson:"folio"`
	SummaryPaidAmount int                `bson:"summaryPaidAmount"`
	ReasonSocial      string             `bson:"reasonSocial"`
	NetAmount         int                `bson:"netAmount"`
	Iva               int                `bson:"iva"`
	DocumentDate      primitive.DateTime `bson:"documentDate"`
	CreditPaid        []Quota            `bson:"creditPaid"`
}

type Quota struct {
	Different    int                `bson:"different"`
	Date         primitive.DateTime `bson:"date"`
	ReceptorName string             `bson:"receptorName"`
	Amount       int                `bson:"amount"`
}
