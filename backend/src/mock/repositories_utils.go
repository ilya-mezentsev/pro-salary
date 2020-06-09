package mock

import (
	"config"
	"errors"
	"github.com/jmoiron/sqlx"
	"models"
	"time"
)

const (
	deleteTablesQuery = `
		DROP TYPE IF EXISTS RATE_TYPE CASCADE;
		DROP TABLE IF EXISTS employees CASCADE;
		DROP TABLE IF EXISTS consumptions CASCADE;
	`
	addTablesQuery = `
		CREATE TYPE RATE_TYPE AS ENUM('hourly', 'monthly');

		CREATE TABLE IF NOT EXISTS employees(
			id SERIAL PRIMARY KEY,
			uuid VARCHAR(36) UNIQUE,
			name VARCHAR(100) NOT NULL,
			pay_type VARCHAR(36) NOT NULL,
			rate_type RATE_TYPE NOT NULL,
			rate INTEGER NOT NULL,
			last_pay_date TIMESTAMP NOT NULL,
			worked_hours INTEGER NOT NULL,
			unpaid_hours INTEGER NOT NULL
		);

		CREATE TABLE IF NOT EXISTS consumptions(
			id SERIAL PRIMARY KEY,
			employee_uuid VARCHAR(36) REFERENCES employees(uuid) ON DELETE CASCADE,
			uuid VARCHAR(36) UNIQUE,
			amount INTEGER NOT NULL
		);`

	addEmployeeQuery = `
	INSERT INTO employees(
		uuid, name, pay_type, rate_type, rate, last_pay_date, worked_hours, unpaid_hours)
	VALUES(:uuid, :name, :pay_type, :rate_type, :rate, :last_pay_date, :worked_hours, :unpaid_hours)`

	addConsumptionQuery = `
	INSERT INTO consumptions(uuid, amount, employee_uuid)
	VALUES(:uuid, :amount, :employee_uuid)`

	getEmployeeQuery = `
	SELECT uuid, name, pay_type, rate_type, rate, last_pay_date FROM employees
	WHERE uuid = $1`

	getEmployeeConsumptionsQuery = `
	SELECT uuid, amount, employee_uuid FROM consumptions
	WHERE employee_uuid = $1`
)

const (
	BadEmployeeId      models.ID = ""
	NotFoundEmployeeId models.ID = "not-found"
)

var (
	SomeError = errors.New("some-error")
)

var (
	date = time.Date(2020, 5, 30, 0, 0, 0, 0, time.Local)

	employees = []models.Employee{
		{
			Id:          "some-id-1",
			Name:        "Nick",
			PayType:     "check",
			RateType:    config.GetHourlyRateName(),
			Rate:        500,
			LastPayDate: date,
			WorkedHours: 230,
			UnpaidHours: 10,
		},
		{
			Id:          "some-id-2",
			Name:        "Alex",
			PayType:     "check-to-email",
			RateType:    config.GetHourlyRateName(),
			Rate:        500,
			LastPayDate: date,
			WorkedHours: 230,
			UnpaidHours: 10,
		},
		{
			Id:          "some-id-3",
			Name:        "John",
			PayType:     "salary-to-account",
			RateType:    config.GetMonthlyRateName(),
			Rate:        20000,
			LastPayDate: date,
			WorkedHours: 230,
			UnpaidHours: 10,
		},
	}

	consumptions = []models.Consumption{
		{
			Id:         "some-id-1",
			EmployeeId: "some-id-1",
			Amount:     100,
		},
		{
			Id:         "some-id-2",
			EmployeeId: "some-id-1",
			Amount:     20,
		},
		{
			Id:         "some-id-3",
			EmployeeId: "some-id-2",
			Amount:     150,
		},
	}
)

func GetAllEmployees() []models.Employee {
	return employees
}

func GetAllConsumptions() []models.Consumption {
	return consumptions
}

func InitTables(db *sqlx.DB) {
	DropTables(db)
	exec(db, addTablesQuery)

	tx := db.MustBegin()
	for _, employee := range employees {
		_, err := tx.NamedExec(addEmployeeQuery, employee)
		if err != nil {
			panic(err)
		}
	}

	for _, consumption := range consumptions {
		_, err := tx.NamedExec(addConsumptionQuery, consumption)
		if err != nil {
			panic(err)
		}
	}

	err := tx.Commit()
	if err != nil {
		panic(err)
	}
}

func DropTables(db *sqlx.DB) {
	exec(db, deleteTablesQuery)
}

func exec(db *sqlx.DB, query string) {
	_, err := db.Exec(query)
	if err != nil {
		panic(err)
	}
}

func GetEmployee(db *sqlx.DB, employeeId models.ID) models.Employee {
	var employee models.Employee
	err := db.Get(&employee, getEmployeeQuery, employeeId)
	if err != nil {
		panic(err)
	}

	return employee
}

func GetEmployeeConsumptions(db *sqlx.DB, employeeId models.ID) []models.Consumption {
	var consumptions []models.Consumption
	err := db.Select(&consumptions, getEmployeeConsumptionsQuery, employeeId)
	if err != nil {
		panic(err)
	}

	return consumptions
}
