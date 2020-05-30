package monthly_rate

import (
	"interfaces"
	"models"
	"time"
)

type Employee struct {
	data models.Employee
}

func New(data models.Employee) interfaces.Employee {
	return Employee{data}
}

func (e Employee) IsPayDay(today time.Time) bool {
	return e.todayIsMonthAfterLastSalaryDate(today)
}

func (e Employee) todayIsMonthAfterLastSalaryDate(today time.Time) bool {
	monthAfterLastSalaryDate := e.data.LastPayDate.AddDate(0, 1, 0)

	return today.Equal(monthAfterLastSalaryDate) || today.After(monthAfterLastSalaryDate)
}

func (e Employee) CalculatePayment(_ int) models.Payment {
	return models.Payment(e.data.Rate)
}
