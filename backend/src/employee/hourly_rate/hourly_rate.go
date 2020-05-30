package hourly_rate

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
	return e.todayIsWeekAfterLastSalaryDate(today)
}

func (e Employee) todayIsWeekAfterLastSalaryDate(today time.Time) bool {
	weekAfterLastSalaryDate := e.data.LastPayDate.AddDate(0, 0, 7)

	return today.Equal(weekAfterLastSalaryDate) || today.After(weekAfterLastSalaryDate)
}
