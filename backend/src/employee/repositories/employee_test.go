package repositories

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"interfaces"
	"mock"
	"os"
	utils "test_utils"
	"testing"
)

var (
	db         *sqlx.DB
	repository interfaces.EmployeesRepository
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

func TestRepository_GetAllSuccess(t *testing.T) {
	mock.InitTables(db)
	defer mock.DropTables(db)

	employees, err := repository.GetAll()

	utils.AssertNil(err, t)
	utils.AssertNotNil(employees, t)
	utils.AssertEqual(len(mock.GetAllEmployees()), len(employees), t)
}

func TestRepository_GetAllError(t *testing.T) {
	mock.DropTables(db)

	_, err := repository.GetAll()

	utils.AssertNotNil(err, t)
}
