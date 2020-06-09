package bookkeeping

import (
	"interfaces"
	"time"
	"types"
)

func PerformPayments(
	employeesFactory interfaces.EmployeesFactory,
	paymentFinalizerConstructor interfaces.PaymentFinalizerConstructor,
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
				paymentFinalizerConstructor,
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
	paymentFinalizerConstructor interfaces.PaymentFinalizerConstructor,
	processing types.PaymentProcessing,
) {
	payment := employee.CalculatePayment()
	paymentFinalizer := paymentFinalizerConstructor.GetPaymentFinalizer(employee.GetPaymentType())

	err := paymentFinalizer.Finish(payment)
	if err != nil {
		processing.Error <- err
		return
	}

	processing.Done <- true
	return
}
