package interfaces

import (
	"models"
	"time"
)

type (
	Employee interface {
		IsPayDay(today time.Time) bool
	}

	EmployeesRepository interface {
		GetAll() ([]models.Employee, error)
	}
)
