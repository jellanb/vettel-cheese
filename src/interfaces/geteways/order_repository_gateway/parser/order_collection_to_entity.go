package parser

import (
	"vettel-backend-app/src/domain/entity"
	"vettel-backend-app/src/infrastructure/database/collections"
	"vettel-backend-app/src/interfaces/geteways/products_repository_gateway/parser"
)

func OrderCollectionToEntity(orderCollection collections.Order) entity.Order {
	var entityOrder entity.Order
	entityOrder.OrderNumber = orderCollection.OrderNumber
	entityOrder.SaleAmount = orderCollection.SaleAmount
	entityOrder.Date = orderCollection.Date
	entityOrder.Products = parser.ProductsCollectionsToEntities(orderCollection.Products)
	entityOrder.Status = orderCollection.Status
	entityOrder.SellerName = orderCollection.SellerName
	entityOrder.Payment = OrderPaymentCollectionToEntity(orderCollection)
	entityOrder.Dispatch = DispatchCollectionToEntity(orderCollection.Dispatch)
	return entityOrder
}

func DispatchCollectionToEntity(dispatchCollection collections.Dispatch) entity.Dispatch {
	var dispatchEntity entity.Dispatch
	dispatchEntity.ReasonSocial = dispatchCollection.ReasonSocial
	dispatchEntity.Products = parser.ProductsCollectionToEntity(dispatchCollection.Products)
	dispatchEntity.CreatedDate = dispatchCollection.CreatedDate
	dispatchEntity.Status = dispatchCollection.Status
	dispatchEntity.Amount = calculateAmount(dispatchCollection.Products)
	dispatchEntity.CreditTime = dispatchCollection.CreditTime
	dispatchEntity.DeliveryName = dispatchCollection.DeliveryName
	dispatchEntity.Phone = dispatchCollection.Phone
	dispatchEntity.AdministratorName = dispatchCollection.AdministratorName
	dispatchEntity.Region = dispatchCollection.Region
	dispatchEntity.Commune = dispatchCollection.Commune
	dispatchEntity.AddressNumber = dispatchCollection.AddressNumber
	dispatchEntity.DeliveryDate = dispatchCollection.DeliveryDate
	dispatchEntity.Address = dispatchCollection.Address
	return dispatchEntity
}

func OrderPaymentCollectionToEntity(orderPayment collections.Order) entity.Payment {
	var paymentEntity entity.Payment
	paymentEntity.PaymentDate = orderPayment.Payment.PaymentDate
	paymentEntity.Folio = orderPayment.Payment.Folio
	paymentEntity.CreditTime = orderPayment.Payment.CreditTime
	paymentEntity.DocumentDate = orderPayment.Payment.DocumentDate
	paymentEntity.Id = orderPayment.Payment.Id
	paymentEntity.Iva = orderPayment.Payment.Iva
	paymentEntity.NetAmount = orderPayment.Payment.NetAmount
	paymentEntity.Amount = orderPayment.Payment.Amount
	paymentEntity.ReasonSocial = orderPayment.Payment.ReasonSocial
	paymentEntity.SummaryPaidAmount = orderPayment.Payment.SummaryPaidAmount
	paymentEntity.Status = orderPayment.Payment.Status
	paymentEntity.CreditPaid = OrderPaymentQuoteCollectionToEntity(orderPayment)
	return paymentEntity
}

func OrderPaymentQuoteCollectionToEntity(orderPayment collections.Order) []entity.Quota {
	var orderPaymentQuotes []entity.Quota
	for _, quota := range orderPayment.Payment.CreditPaid {
		var orderPaymentQuote entity.Quota
		orderPaymentQuote.Date = quota.Date
		orderPaymentQuote.Different = quota.Different
		orderPaymentQuote.ReceptorName = quota.ReceptorName
		orderPaymentQuote.Amount = quota.Amount
		orderPaymentQuotes = append(orderPaymentQuotes, orderPaymentQuote)
	}
	return orderPaymentQuotes
}

func calculateAmount(productsCollection []collections.Product) int {
	totalAmount := 0
	for _, product := range productsCollection {
		totalAmount += product.SaleAmount
	}
	return totalAmount
}
