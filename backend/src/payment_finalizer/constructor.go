package payment_finalizer

import (
	"interfaces"
	"models"
	"payment_finalizer/finalizer"
)

type Constructor struct {
	repository              interfaces.PaymentFinalizerRepository
	payTypeToCheckProcessor map[models.PayType]interfaces.CheckProcessor
	defaultCheckProcessor   interfaces.CheckProcessor
}

func New(repository interfaces.PaymentFinalizerRepository) interfaces.PaymentFinalizerConstructor {
	return &Constructor{
		repository:              repository,
		payTypeToCheckProcessor: map[models.PayType]interfaces.CheckProcessor{},
		defaultCheckProcessor:   DefaultCheckProcessor{},
	}
}

func (c Constructor) GetPaymentFinalizer(payType models.PayType) interfaces.PaymentFinalizer {
	employeeCheckProcessor, found := c.payTypeToCheckProcessor[payType]
	if !found {
		employeeCheckProcessor = c.defaultCheckProcessor
	}

	return finalizer.New(c.repository, employeeCheckProcessor)
}

func (c *Constructor) SetCheckProcessor(payType models.PayType, processor interfaces.CheckProcessor) {
	c.payTypeToCheckProcessor[payType] = processor
}

func (c *Constructor) SetDefaultCheckProcessor(processor interfaces.CheckProcessor) {
	c.defaultCheckProcessor = processor
}

type DefaultCheckProcessor struct {
}

func (p DefaultCheckProcessor) Process(models.Check) error {
	return nil
}
