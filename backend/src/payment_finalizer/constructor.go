package payment_finalizer

import (
	"interfaces"
	"models"
)

type Constructor struct {
}

func (c Constructor) GetPaymentFinalizer(payType models.PayType) (interfaces.PaymentFinalizer, error) {
	panic("implement me")
}

type DefaultCheckProcessor struct {
}

func (p DefaultCheckProcessor) Process(_ models.Check) error {
	return nil
}
