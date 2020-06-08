package repositories

import (
	"app_internals"
	"github.com/jmoiron/sqlx"
	"interfaces"
	"models"
)

const (
	addUnpaidToWorkedHoursQuery = `
	UPDATE employees
	SET worked_hours = worked_hours + unpaid_hours
	WHERE uuid = $1`

	setUnpaidHoursToZeroQuery = `
	UPDATE employees
	SET unpaid_hours = 0
	WHERE uuid = $1`

	getEmployeeConsumptionsQuery = `
	SELECT uuid, amount, employee_uuid FROM consumptions
	WHERE employee_uuid = $1`

	completeConsumptionsQuery = `DELETE FROM consumptions WHERE employee_uuid = $1`
)

type Repository struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) interfaces.PaymentFinalizerRepository {
	return Repository{db}
}

func (r Repository) AddUnpaidToWorkedHours(employeeId models.ID) error {
	res, err := r.db.Exec(addUnpaidToWorkedHoursQuery, employeeId)
	if err != nil {
		return err
	}

	affectedRows, err := res.RowsAffected()
	if err != nil {
		return err
	} else if affectedRows == 0 {
		return app_internals.EmployeeNotFound
	}

	return nil
}

func (r Repository) SetUnpaidHoursToZero(employeeId models.ID) error {
	res, err := r.db.Exec(setUnpaidHoursToZeroQuery, employeeId)
	if err != nil {
		return err
	}

	affectedRows, err := res.RowsAffected()
	if err != nil {
		return err
	} else if affectedRows == 0 {
		return app_internals.EmployeeNotFound
	}

	return nil
}

func (r Repository) GetEmployeeConsumptions(employeeId models.ID) ([]models.Consumption, error) {
	var consumptions []models.Consumption
	err := r.db.Select(&consumptions, getEmployeeConsumptionsQuery, employeeId)
	if err != nil {
		return nil, err
	}

	return consumptions, nil
}

func (r Repository) MakeConsumptionsCompleted(employeeId models.ID) error {
	res, err := r.db.Exec(completeConsumptionsQuery, employeeId)
	if err != nil {
		return err
	}

	affectedRows, err := res.RowsAffected()
	if err != nil {
		return err
	} else if affectedRows == 0 {
		return app_internals.EmployeeNotFound
	}

	return nil
}
