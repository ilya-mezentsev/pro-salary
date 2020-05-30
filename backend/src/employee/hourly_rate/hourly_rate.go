package hourly_rate

import (
	"interfaces"
	"models"
	"time"
)

const (
	weekWorkHours         = 40
	extraHoursCoefficient = 1.5
)

type Employee struct {
	data models.Employee
}

func New(data models.Employee) interfaces.Employee {
	return Employee{data}
}

func (e Employee) IsPayDay(today time.Time) bool {
	return e.todayIsWeekAfterLastSalaryDate(today)
}

func (e Employee) todayIsWeekAfterLastSalaryDate(today time.Time) bool {
	weekAfterLastSalaryDate := e.data.LastPayDate.AddDate(0, 0, 7)

	return today.Equal(weekAfterLastSalaryDate) || today.After(weekAfterLastSalaryDate)
}

func (e Employee) CalculatePayment(workedHours int) models.Payment {
	if workedHours > weekWorkHours {
		return models.Payment(
			e.data.Rate*weekWorkHours +
				e.data.Rate*models.Rate(extraHoursCoefficient)*models.Rate(workedHours%weekWorkHours))
	} else {
		return models.Payment(e.data.Rate * models.Rate(workedHours))
	}
}
