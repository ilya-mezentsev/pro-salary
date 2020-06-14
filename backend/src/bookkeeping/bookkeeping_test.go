package bookkeeping

import (
	employeeFactory "employee"
	employeeRepository "employee/repositories"
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"interfaces"
	"mock"
	"mock/bookkeeping"
	"models"
	"os"
	"payment_finalizer"
	paymentFinalizerRepository "payment_finalizer/repositories"
	utils "test_utils"
	"testing"
	"time"
)

var (
	db                          *sqlx.DB
	employeesFactory            interfaces.EmployeesFactory
	paymentFinalizerConstructor interfaces.PaymentFinalizerConstructor

	savingChecksProcessor = bookkeeping.SaveChecksProcessor{}
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

	employeesFactory = employeeFactory.New(employeeRepository.New(db))
	paymentFinalizerConstructor = payment_finalizer.New(paymentFinalizerRepository.New(db))
}

func getConsumptionsAmount(consumptions []models.Consumption) models.PaymentAmount {
	var sum models.PaymentAmount
	for _, consumption := range consumptions {
		sum += consumption.Amount
	}

	return sum
}

func TestPerformPaymentsSimpleCallSuccess(t *testing.T) {
	mock.InitTables(db)
	defer mock.DropTables(db)

	err := PerformPayments(
		employeesFactory,
		paymentFinalizerConstructor,
		time.Now(),
	)

	utils.AssertNil(err, t)
}

func TestPerformPaymentsLookupChecks(t *testing.T) {
	mock.InitTables(db)
	defer mock.DropTables(db)
	paymentFinalizerConstructor.SetDefaultCheckProcessor(&savingChecksProcessor)
	now := time.Now()
	currentYear, currentMonth, currentDay := now.Date()

	err := PerformPayments(
		employeesFactory,
		paymentFinalizerConstructor,
		now,
	)

	utils.AssertNil(err, t)
	for _, check := range savingChecksProcessor.GetSavedChecks() {
		utils.AssertEqual(check.Total, check.Amount-getConsumptionsAmount(check.Consumptions), t)

		utils.AssertEqual(0, len(mock.GetEmployeeConsumptions(db, check.EmployeeId)), t)
		year, month, day := mock.GetEmployee(db, check.EmployeeId).LastPayDate.Date()
		utils.AssertEqual(currentYear, year, t)
		utils.AssertEqual(currentMonth, month, t)
		utils.AssertEqual(currentDay, day, t)
	}
}

func TestPerformPaymentsBadPaymentFinalizer(t *testing.T) {
	mock.InitTables(db)
	defer mock.DropTables(db)

	err := PerformPayments(
		employeesFactory,
		bookkeeping.BadPaymentFinalizerConstructor{},
		time.Now(),
	)

	utils.AssertNotNil(err, t)
}

func TestPerformPaymentsSomeError(t *testing.T) {
	mock.DropTables(db)

	err := PerformPayments(
		employeesFactory,
		paymentFinalizerConstructor,
		time.Now(),
	)

	utils.AssertNotNil(err, t)
}
