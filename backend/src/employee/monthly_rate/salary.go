package monthly_rate

import (
	"models"
	"time"
)

type salary struct {
	data models.Employee
}

func (s salary) isPayDay(today time.Time) bool {
	return s.todayIsMonthAfterLastSalaryDate(today)
}

func (s salary) todayIsMonthAfterLastSalaryDate(today time.Time) bool {
	monthAfterLastSalaryDate := s.data.LastPayDate.AddDate(0, 1, 0)

	return today.Equal(monthAfterLastSalaryDate) || today.After(monthAfterLastSalaryDate)
}
