package hourly_rate

import (
	"models"
	"time"
)

const weekWorkHours = 40

type salary struct {
	data models.Employee
}

func (s salary) isPayDay(today time.Time) bool {
	return s.todayIsWeekAfterLastSalaryDate(today)
}

func (s salary) todayIsWeekAfterLastSalaryDate(today time.Time) bool {
	weekAfterLastSalaryDate := s.data.LastPayDate.AddDate(0, 0, 7)

	return today.Equal(weekAfterLastSalaryDate) || today.After(weekAfterLastSalaryDate)
}
