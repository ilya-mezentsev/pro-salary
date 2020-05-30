package monthly_rate

import (
	"interfaces"
	"models"
	"time"
)

type Employee struct {
	data   models.Employee
	salary salary
}

func New(data models.Employee) interfaces.Employee {
	return Employee{
		data:   data,
		salary: salary{data},
	}
}

func (e Employee) IsPayDay(today time.Time) bool {
	return e.salary.isPayDay(today)
}

func (e Employee) CalculatePayment(_ int) models.Payment {
	return models.Payment(e.data.Rate)
}

func (e Employee) GetPaymentType() models.PayType {
	return e.data.PayType
}
