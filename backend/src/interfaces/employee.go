package interfaces

import (
	"models"
	"time"
)

type (
	Employee interface {
		IsPayDay(today time.Time) bool
		CalculatePayment(workedHours int) models.Payment
		GetPaymentType() models.PayType
	}

	EmployeesRepository interface {
		GetAll() ([]models.Employee, error)
	}

	EmployeesFactory interface {
		GetAll() ([]Employee, error)
	}
)
