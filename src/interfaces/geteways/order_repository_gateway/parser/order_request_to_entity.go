package parser

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"vettel-backend-app/src/domain/entity"
	"vettel-backend-app/src/infrastructure/web_server/models/request"
)

func OrderRequestToEntity(orderRequest request.Order) entity.Order {
	var orderEntity entity.Order
	orderEntity.OrderNumber = orderRequest.OrderNumber
	orderEntity.Date = orderRequest.Date
	orderEntity.SaleAmount = orderRequest.SaleAmount
	orderEntity.Status = orderRequest.Status
	orderEntity.ReasonSocial = orderRequest.ReasonSocial
	orderEntity.SellerName = orderRequest.SellerName
	orderEntity.Products = OrderRequestToOrderEntityProducts(orderRequest)
	orderEntity.Dispatch = OrderRequestToDispatchEntity(orderRequest)
	orderEntity.Payment = OrderRequestToPaymentEntity(orderRequest)

	return orderEntity
}

func OrderRequestToOrderEntityProducts(order request.Order) []entity.Product {
	var productsEntity []entity.Product
	var productEntity entity.Product
	for _, product := range order.Products {
		productEntity.Quantity = product.Quantity
		productEntity.Barcode = product.Barcode
		productEntity.SaleAmount = product.SaleAmount
		productEntity.Description = product.Description
		productEntity.CreatedDate = primitive.NewDateTimeFromTime(product.CreatedDate)
		productEntity.UpdateDate = primitive.NewDateTimeFromTime(product.UpdateDate)
		productsEntity = append(productsEntity, productEntity)
	}
	return productsEntity
}

func OrderRequestToDispatchEntity(order request.Order) entity.Dispatch {
	var dispatchEntity entity.Dispatch
	dispatchEntity.Amount = order.Payment.Amount
	dispatchEntity.Status = order.Dispatch.Status
	dispatchEntity.CreditTime = order.Dispatch.CreditTime
	dispatchEntity.Products = OrderRequestToOrderEntityProducts(order)
	dispatchEntity.ReasonSocial = order.Dispatch.ReasonSocial
	dispatchEntity.DeliveryDate = order.Dispatch.DeliveryDate
	dispatchEntity.CreatedDate = order.Dispatch.CreatedDate
	dispatchEntity.AdministratorName = order.Dispatch.AdministratorName
	dispatchEntity.DeliveryName = order.Dispatch.DeliveryName
	dispatchEntity.Phone = order.Dispatch.Phone
	dispatchEntity.Region = order.Dispatch.Region
	dispatchEntity.Commune = order.Dispatch.Commune
	dispatchEntity.AddressNumber = order.Dispatch.AddressNumber
	dispatchEntity.Address = order.Dispatch.Address
	return dispatchEntity
}

func OrderRequestToPaymentEntity(order request.Order) entity.Payment {
	var paymentEntity entity.Payment
	paymentEntity.Id = order.Payment.Id
	paymentEntity.PaymentDate = order.Payment.PaymentDate
	paymentEntity.Status = order.Payment.Status
	paymentEntity.Amount = order.Payment.Amount
	paymentEntity.ReasonSocial = order.Payment.ReasonSocial
	paymentEntity.Iva = order.Payment.Iva
	paymentEntity.NetAmount = order.Payment.NetAmount
	paymentEntity.SummaryPaidAmount = order.Payment.SummaryPaidAmount
	paymentEntity.Folio = order.Payment.Folio
	paymentEntity.CreditTime = order.Payment.CreditTime
	paymentEntity.DocumentDate = order.Payment.DocumentDate
	paymentEntity.CreditPaid = GenerateCreditPaidEntity(order)
	return paymentEntity
}

func GenerateCreditPaidEntity(order request.Order) []entity.Quota {
	var quoteEntities []entity.Quota
	for _, quota := range order.Payment.CreditPaid {
		var newQuote entity.Quota
		newQuote.Date = quota.Date
		newQuote.Different = quota.Different
		newQuote.ReceptorName = quota.ReceptorName
		newQuote.Amount = quota.Amount
		quoteEntities = append(quoteEntities, newQuote)
	}
	return quoteEntities
}
