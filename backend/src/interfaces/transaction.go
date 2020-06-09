package interfaces

import "models"

type (
	PaymentFinalizerConstructor interface {
		GetPaymentFinalizer(payType models.PayType) PaymentFinalizer
		SetCheckProcessor(payType models.PayType, processor CheckProcessor)
		SetDefaultCheckProcessor(processor CheckProcessor)
	}

	PaymentFinalizer interface {
		Finish(payment models.Payment) error
	}

	CheckProcessor interface {
		Process(check models.Check) error
	}

	PaymentFinalizerRepository interface {
		AddUnpaidToWorkedHours(employeeId models.ID) error
		SetUnpaidHoursToZero(employeeId models.ID) error
		GetEmployeeConsumptions(employeeId models.ID) ([]models.Consumption, error)
		MakeConsumptionsCompleted(employeeId models.ID) error
	}
)
