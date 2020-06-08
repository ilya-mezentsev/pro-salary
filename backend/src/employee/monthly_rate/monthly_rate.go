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

func (e Employee) CalculatePayment() models.Payment {
	return models.Payment{
		EmployeeId: e.data.Id,
		Amount:     models.PaymentAmount(e.data.Rate),
	}
}

func (e Employee) GetPaymentType() models.PayType {
	return e.data.PayType
}
