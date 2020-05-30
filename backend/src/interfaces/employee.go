package interfaces

import (
	"models"
	"time"
)

type (
	Employee interface {
		IsPayDay(today time.Time) bool
		CalculatePayment(workedHours int) models.Payment
	}

	EmployeesRepository interface {
		GetAll() ([]models.Employee, error)
	}
)
