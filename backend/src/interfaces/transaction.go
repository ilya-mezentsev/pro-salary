package interfaces

import "models"

type (
	PaymentResultFactory interface {
		GetPaymentResult(payType models.PayType) (PaymentResult, error)
	}

	PaymentResult interface {
		FinishTransaction(payment models.Payment) error
	}
)
