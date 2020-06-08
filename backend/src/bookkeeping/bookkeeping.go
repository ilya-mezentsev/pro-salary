package bookkeeping

import (
	"interfaces"
	"time"
	"types"
)

func PerformPayments(
	employeesFactory interfaces.EmployeesFactory,
	paymentResultFactory interfaces.PaymentFinalizerConstructor,
	today time.Time,
) error {
	employees, err := employeesFactory.GetAll()
	if err != nil {
		return err
	}

	processedPaymentsCount := 0
	paymentsCount := 0
	processing := types.PaymentProcessing{
		Error: make(chan error),
		Done:  make(chan bool),
	}
	for _, employee := range employees {
		if employee.IsPayDay(today) {
			paymentsCount++
			go performPayment(
				employee,
				paymentResultFactory,
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
			if processedPaymentsCount >= paymentsCount {
				return nil
			}
		}
	}
}

func performPayment(
	employee interfaces.Employee,
	paymentResultFactory interfaces.PaymentFinalizerConstructor,
	processing types.PaymentProcessing,
) {
	payment := employee.CalculatePayment()
	paymentFinalizer, err := paymentResultFactory.GetPaymentFinalizer(employee.GetPaymentType())
	if err != nil {
		processing.Error <- err
		return
	}

	err = paymentFinalizer.Finish(payment)
	if err != nil {
		processing.Error <- err
		return
	}

	processing.Done <- true
	return
}
