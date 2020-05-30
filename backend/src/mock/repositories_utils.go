package mock

import (
	"config"
	"github.com/jmoiron/sqlx"
	"models"
	"time"
)

const (
	deleteTablesQuery = `
		DROP TYPE IF EXISTS RATE_TYPE CASCADE;
		DROP TABLE IF EXISTS employees CASCADE;
	`
	addTablesQuery = `
		CREATE TYPE RATE_TYPE AS ENUM('hourly', 'monthly');

		CREATE TABLE IF NOT EXISTS employees(
			id SERIAL PRIMARY KEY,
			uuid VARCHAR(36) NOT NULL,
			name VARCHAR(100) NOT NULL,
			pay_type VARCHAR(36) NOT NULL,
			rate_type RATE_TYPE NOT NULL,
			rate INTEGER NOT NULL,
			last_pay_date TIMESTAMP NOT NULL
		);`
	addEmployeeQuery = `
	INSERT INTO employees(uuid, name, pay_type, rate_type, rate, last_pay_date)
	VALUES(:uuid, :name, :pay_type, :rate_type, :rate, :last_pay_date)`
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
		},
		{
			Id:          "some-id-2",
			Name:        "Alex",
			PayType:     "check-to-email",
			RateType:    config.GetHourlyRateName(),
			Rate:        500,
			LastPayDate: date,
		},
		{
			Id:          "some-id-3",
			Name:        "John",
			PayType:     "salary-to-account",
			RateType:    config.GetMonthlyRateName(),
			Rate:        20000,
			LastPayDate: date,
		},
	}
)

func GetAllEmployees() []models.Employee {
	return employees
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
