package hourly_rate

import (
	"interfaces"
	"models"
	"time"
)

type Employee struct {
	data    models.Employee
	salary  salary
	payment payment
}

func New(data models.Employee) interfaces.Employee {
	return Employee{
		data:    data,
		salary:  salary{data},
		payment: payment{data},
	}
}

func (e Employee) IsPayDay(today time.Time) bool {
	return e.salary.isPayDay(today)
}

func (e Employee) CalculatePayment(workedHours int) models.Payment {
	return e.payment.calculate(workedHours)
}

func (e Employee) GetPaymentType() models.PayType {
	return e.data.PayType
}
