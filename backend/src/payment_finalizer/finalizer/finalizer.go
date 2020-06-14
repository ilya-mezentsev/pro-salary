package finalizer

import (
	"interfaces"
	"models"
)

type Finalizer struct {
	repository             interfaces.PaymentFinalizerRepository
	employeeCheckProcessor interfaces.CheckProcessor
}

func New(
	repository interfaces.PaymentFinalizerRepository,
	employeeCheckProcessor interfaces.CheckProcessor,
) interfaces.PaymentFinalizer {
	return Finalizer{repository, employeeCheckProcessor}
}

func (f Finalizer) Finish(payment models.Payment) error {
	err := f.repository.AddUnpaidToWorkedHours(payment.EmployeeId)
	if err != nil {
		return err
	}

	err = f.repository.SetUnpaidHoursToZeroAndResetLastPayDate(payment.EmployeeId)
	if err != nil {
		return err
	}

	check, err := f.createCheck(payment)
	if err != nil {
		return err
	}

	err = f.repository.MakeConsumptionsCompleted(check.EmployeeId)
	if err != nil {
		return err
	}

	return f.employeeCheckProcessor.Process(check)
}

func (f Finalizer) createCheck(payment models.Payment) (models.Check, error) {
	consumptions, err := f.repository.GetEmployeeConsumptions(payment.EmployeeId)
	if err != nil {
		return models.Check{}, err
	}

	total := payment.Amount
	for _, consumption := range consumptions {
		total -= consumption.Amount
	}

	return models.Check{
		EmployeeId:   payment.EmployeeId,
		Amount:       payment.Amount,
		Consumptions: consumptions,
		Total:        total,
	}, nil
}
