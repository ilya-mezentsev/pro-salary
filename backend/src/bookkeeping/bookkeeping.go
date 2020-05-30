package bookkeeping

import (
	"interfaces"
	"time"
	"types"
)

func PerformPayments(
	employeesFactory interfaces.EmployeesFactory,
	paymentResultFactory interfaces.PaymentResultFactory,
	workedHours int,
	today time.Time,
) error {
	employees, err := employeesFactory.GetAll()
	if err != nil {
		return err
	}

	processedPaymentsCount := 0
	employeesCount := len(employees)
	processing := types.PaymentProcessing{
		Error: make(chan error),
		Done:  make(chan bool),
	}
	for _, employee := range employees {
		if employee.IsPayDay(today) {
			go performPayment(
				employee,
				paymentResultFactory,
				workedHours,
				processing,
			)
		}
	}

	for {
		select {
		case <-processing.Done:
			processedPaymentsCount++
		case err = <-processing.Error:
			return err
		default:
			if processedPaymentsCount >= employeesCount {
				return nil
			}
		}
	}
}

func performPayment(
	employee interfaces.Employee,
	paymentResultFactory interfaces.PaymentResultFactory,
	workedHours int,
	processing types.PaymentProcessing,
) {
	payment := employee.CalculatePayment(workedHours)
	paymentResult, err := paymentResultFactory.GetPaymentResult(employee.GetPaymentType())
	if err != nil {
		processing.Error <- err
		return
	}

	err = paymentResult.FinishTransaction(payment)
	if err != nil {
		processing.Error <- err
		return
	}

	processing.Done <- true
	return
}
