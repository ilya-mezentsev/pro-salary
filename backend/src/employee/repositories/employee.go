package repositories

import (
	"github.com/jmoiron/sqlx"
	"interfaces"
	"models"
)

const (
	selectEmployeesQuery = `
	SELECT uuid, name, pay_type, rate_type, rate, last_pay_date, worked_hours, unpaid_hours
	FROM employees`
)

type Repository struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) interfaces.EmployeesRepository {
	return Repository{db}
}

func (r Repository) GetAll() ([]models.Employee, error) {
	var employees []models.Employee
	err := r.db.Select(&employees, selectEmployeesQuery)

	return employees, err
}
