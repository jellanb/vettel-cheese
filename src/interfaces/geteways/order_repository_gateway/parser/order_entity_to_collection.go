package parser

import (
	"vettel-backend-app/src/domain/entity"
	"vettel-backend-app/src/infrastructure/database/collections"
	parser_dispatch "vettel-backend-app/src/interfaces/geteways/dispatch_repository_gateway/parser"
	"vettel-backend-app/src/interfaces/geteways/products_repository_gateway/parser"
)

func OrderEntityToCollection(orderEntity entity.Order) collections.Order {
	var orderCollection collections.Order
	orderCollection.OrderNumber = orderEntity.OrderNumber
	orderCollection.SaleAmount = orderEntity.SaleAmount
	orderCollection.Date = orderEntity.Date
	orderCollection.Products = parser.ProductsEntitiesToCollection(orderEntity.Products)
	orderCollection.Status = orderEntity.Status
	orderCollection.SellerName = orderEntity.SellerName
	orderCollection.Payment = OrderEntityToOrderPaymentCollection(orderEntity)
	orderCollection.Dispatch = parser_dispatch.DispatchEntityToCollection(orderEntity.Dispatch)
	orderCollection.ReasonSocial = orderEntity.ReasonSocial
	return orderCollection
}

func OrderEntityToOrderPaymentCollection(order entity.Order) collections.Payment {
	var paymentCollection collections.Payment
	paymentCollection.PaymentDate = order.Payment.PaymentDate
	paymentCollection.Folio = order.Payment.Folio
	paymentCollection.CreditTime = order.Payment.CreditTime
	paymentCollection.DocumentDate = order.Payment.DocumentDate
	paymentCollection.Iva = order.Payment.Iva
	paymentCollection.NetAmount = order.Payment.NetAmount
	paymentCollection.Id = order.Payment.Id
	paymentCollection.Amount = order.Payment.Amount
	paymentCollection.ReasonSocial = order.Payment.ReasonSocial
	paymentCollection.SummaryPaidAmount = order.Payment.SummaryPaidAmount
	paymentCollection.Status = order.Payment.Status
	paymentCollection.CreditPaid = OrderEntityPaymentToQuoteCollection(order)
	return paymentCollection
}

func OrderEntityPaymentToQuoteCollection(order entity.Order) []collections.Quota {
	var orderPaymentQuotesCollection []collections.Quota
	for _, quota := range order.Payment.CreditPaid {
		var orderPaymentQuoteColl collections.Quota
		orderPaymentQuoteColl.Date = quota.Date
		orderPaymentQuoteColl.Different = quota.Different
		orderPaymentQuoteColl.ReceptorName = quota.ReceptorName
		orderPaymentQuoteColl.Amount = quota.Amount
		orderPaymentQuotesCollection = append(orderPaymentQuotesCollection, orderPaymentQuoteColl)
	}
	return orderPaymentQuotesCollection
}
