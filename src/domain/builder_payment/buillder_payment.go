package builder_payment

import "vettel-backend-app/src/domain/entity"

type BuilderPaymentInterface interface {
	BuilderNewPayment() entity.Payment
}

type BuilderPayment struct {
	payment entity.Payment
}

func NewBuilderPayment() BuilderPaymentInterface {
	return &BuilderPayment{}
}

func (p BuilderPayment) BuilderNewPayment() entity.Payment {
	return p.payment
}
