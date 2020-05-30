package employee

import (
	"config"
	"employee/hourly_rate"
	"employee/monthly_rate"
	"errors"
	"mock"
	"mock/employee"
	utils "test_utils"
	"testing"
)

var (
	repository = employee.RepositoryMock{}
	factory    = New(repository)
)

func TestFactory_GetAllSuccess(t *testing.T) {
	employees, err := factory.GetAll()

	utils.AssertNil(err, t)
	utils.AssertNotNil(employees, t)
}

func TestFactory_GetAllCorrectTypes(t *testing.T) {
	employees, _ := factory.GetAll()
	employeesData := mock.GetAllEmployees()

	utils.AssertEqual(len(employeesData), len(employees), t)
	for employeeIndex, e := range employees {
		if employeesData[employeeIndex].RateType == config.GetHourlyRateName() {
			_, typeCorrect := e.(hourly_rate.Employee)
			utils.AssertTrue(typeCorrect, t)
		} else if employeesData[employeeIndex].RateType == config.GetMonthlyRateName() {
			_, typeCorrect := e.(monthly_rate.Employee)
			utils.AssertTrue(typeCorrect, t)
		}
	}
}

func TestFactory_GetAllSomeError(t *testing.T) {
	e := errors.New("some error")
	repository.Error = e
	factory = New(repository)
	defer func() {
		repository.Error = nil
	}()

	_, err := factory.GetAll()

	utils.AssertErrorsEqual(e, err, t)
}
