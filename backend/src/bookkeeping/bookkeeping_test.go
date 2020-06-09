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

	err := PerformPayments(
		employeesFactory,
		paymentFinalizerConstructor,
		time.Now(),
	)

	utils.AssertNil(err, t)
}
