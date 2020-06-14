package repositories

import (
	"app_internals"
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"interfaces"
	"mock"
	"os"
	utils "test_utils"
	"testing"
	"time"
)

var (
	db         *sqlx.DB
	repository interfaces.PaymentFinalizerRepository
)

func init() {
	connStr := os.Getenv("CONN_STR")
	if connStr == "" {
		fmt.Println("CONN_STR is not provided")
		os.Exit(1)
	}

	var err error
	db, err = sqlx.Open("postgres", connStr)
	if err != nil {
		fmt.Printf("Error while opening DB: %v\n", err)
		os.Exit(1)
	}

	repository = New(db)
}

func TestRepository_AddUnpaidToWorkedHoursSuccess(t *testing.T) {
	mock.InitTables(db)
	defer mock.DropTables(db)
	employee := mock.GetEmployee(db, mock.GetAllEmployees()[0].Id)
	expectedWorkedHours := employee.UnpaidHours + employee.WorkedHours

	err := repository.AddUnpaidToWorkedHours(mock.GetAllEmployees()[0].Id)

	utils.AssertNil(err, t)
	utils.AssertEqual(expectedWorkedHours, employee.WorkedHours, t)
}

func TestRepository_AddUnpaidToWorkedHoursNotFoundEmployeeId(t *testing.T) {
	mock.InitTables(db)
	defer mock.DropTables(db)

	err := repository.AddUnpaidToWorkedHours(mock.NotFoundEmployeeId)

	utils.AssertErrorsEqual(app_internals.EmployeeNotFound, err, t)
}

func TestRepository_AddUnpaidToWorkedHoursSomeError(t *testing.T) {
	mock.DropTables(db)

	err := repository.AddUnpaidToWorkedHours(mock.GetAllEmployees()[0].Id)

	utils.AssertNotNil(err, t)
}

func TestRepository_SetUnpaidHoursToZeroSuccess(t *testing.T) {
	mock.InitTables(db)
	defer mock.DropTables(db)
	currentYear, currentMonth, currentDay := time.Now().Date()

	err := repository.SetUnpaidHoursToZeroAndResetLastPayDate(mock.GetAllEmployees()[0].Id)
	employee := mock.GetEmployee(db, mock.GetAllEmployees()[0].Id)

	utils.AssertNil(err, t)
	utils.AssertEqual(0, employee.UnpaidHours, t)
	year, month, day := employee.LastPayDate.Date()
	utils.AssertEqual(currentYear, year, t)
	utils.AssertEqual(currentMonth, month, t)
	utils.AssertEqual(currentDay, day, t)
}

func TestRepository_SetUnpaidHoursToZeroNotFoundEmployeeId(t *testing.T) {
	mock.InitTables(db)
	defer mock.DropTables(db)

	err := repository.SetUnpaidHoursToZeroAndResetLastPayDate(mock.NotFoundEmployeeId)

	utils.AssertErrorsEqual(app_internals.EmployeeNotFound, err, t)
}

func TestRepository_SetUnpaidHoursToZeroSomeError(t *testing.T) {
	mock.DropTables(db)

	err := repository.SetUnpaidHoursToZeroAndResetLastPayDate(mock.GetAllEmployees()[0].Id)

	utils.AssertNotNil(err, t)
}

func TestRepository_GetEmployeeConsumptionsSuccess(t *testing.T) {
	mock.InitTables(db)
	defer mock.DropTables(db)
	expectedConsumptions := mock.GetEmployeeConsumptions(db, mock.GetAllEmployees()[0].Id)

	consumptions, err := repository.GetEmployeeConsumptions(mock.GetAllEmployees()[0].Id)

	utils.AssertNil(err, t)
	utils.AssertEqual(len(expectedConsumptions), len(consumptions), t)
	for consumptionId, consumption := range consumptions {
		utils.AssertEqual(expectedConsumptions[consumptionId], consumption, t)
	}
}

func TestRepository_GetEmployeeConsumptionsSomeError(t *testing.T) {
	mock.DropTables(db)

	_, err := repository.GetEmployeeConsumptions(mock.GetAllEmployees()[0].Id)

	utils.AssertNotNil(err, t)
}

func TestRepository_MakeConsumptionsCompletedSuccess(t *testing.T) {
	mock.InitTables(db)
	defer mock.DropTables(db)

	err := repository.MakeConsumptionsCompleted(mock.GetAllEmployees()[0].Id)

	utils.AssertNil(err, t)
	utils.AssertEqual(
		0,
		len(mock.GetEmployeeConsumptions(db, mock.GetAllEmployees()[0].Id)),
		t,
	)
}

func TestRepository_MakeConsumptionsCompletedNotFoundEmployeeId(t *testing.T) {
	mock.InitTables(db)
	defer mock.DropTables(db)

	err := repository.MakeConsumptionsCompleted(mock.NotFoundEmployeeId)

	utils.AssertErrorsEqual(app_internals.EmployeeNotFound, err, t)
}

func TestRepository_MakeConsumptionsCompletedSomeError(t *testing.T) {
	mock.DropTables(db)

	err := repository.MakeConsumptionsCompleted(mock.GetAllEmployees()[0].Id)

	utils.AssertNotNil(err, t)
}
