package request

import "go.mongodb.org/mongo-driver/bson/primitive"

type Order struct {
	OrderNumber  int                `json:"orderNumber"`
	Date         primitive.DateTime `json:"date"`
	SellerName   string             `json:"sellerName"`
	Status       string             `json:"status"`
	SaleAmount   int                `json:"saleAmount"`
	Products     []Product          `json:"products"`
	Dispatch     OrderDispatch      `json:"dispatch"`
	Payment      Payment            `json:"payment"`
	ReasonSocial string             `json:"reasonSocial"`
}

type OrderDispatch struct {
	Id                int                `json:"id"`
	CreatedDate       primitive.DateTime `json:"createdDate"`
	ReasonSocial      string             `json:"reasonSocial"`
	Address           string             `json:"address"`
	Commune           string             `json:"commune"`
	AddressNumber     int                `json:"addressNumber"`
	Region            string             `json:"region"`
	AdministratorName string             `json:"administratorName"`
	Phone             string             `json:"phone"`
	DeliveryDate      primitive.DateTime `json:"deliveryDate"`
	Status            string             `json:"status"`
	DeliveryName      string             `json:"deliveryName"`
	Products          []Product          `json:"products"`
	CreditTime        primitive.DateTime `json:"creditTime"`
}

type Payment struct {
	Id                int                `json:"id"`
	PaymentDate       primitive.DateTime `json:"paymentDate"`
	CreditTime        primitive.DateTime `json:"creditTime"`
	Amount            int                `json:"amount"`
	Status            string             `json:"status"`
	Folio             string             `json:"folio"`
	SummaryPaidAmount int                `json:"summaryPaidAmount"`
	ReasonSocial      string             `json:"reasonSocial"`
	NetAmount         int                `json:"netAmount"`
	Iva               int                `json:"iva"`
	DocumentDate      primitive.DateTime `json:"documentDate"`
	CreditPaid        []Quota            `json:"creditPaid"`
}
type Quota struct {
	Different    int                `json:"different"`
	Date         primitive.DateTime `json:"date"`
	ReceptorName string             `json:"receptorName"`
	Amount       int                `json:amount`
}
