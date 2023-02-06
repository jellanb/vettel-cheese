package parser

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
	"vettel-backend-app/src/domain/builder_order"
	"vettel-backend-app/src/domain/builder_payment"
	"vettel-backend-app/src/domain/entity"
	"vettel-backend-app/src/domain/status_domain/order_status"
	"vettel-backend-app/src/domain/status_domain/payment_status"
)

func GenerateOrderEntityFromDispatch(dispatch entity.Dispatch) entity.Order {
	orderBuilder := builder_order.NewOrderBuilder()
	builderPayment := builder_payment.NewBuilderPayment()
	order := orderBuilder.BuildNewOrder()
	payment := builderPayment.BuilderNewPayment()

	order.Status = order_status.Created
	order.Date = primitive.NewDateTimeFromTime(time.Now())
	order.SellerName = dispatch.DeliveryName
	order.SaleAmount = dispatch.Amount
	order.Dispatch = dispatch
	order.Products = dispatch.Products
	order.ReasonSocial = dispatch.ReasonSocial

	payment.PaymentDate = dispatch.CreditTime
	payment.Status = payment_status.Pending
	payment.ReasonSocial = dispatch.ReasonSocial
	payment.CreditTime = dispatch.CreditTime
	payment.Amount = dispatch.Amount
	payment.Folio = ""
	payment.SummaryPaidAmount = 0

	order.Payment = payment

	return order
}
