package payment_finalizer

import (
	"app_internals"
	"mock"
	"models"
)

type RepositoryMock struct {
}

func (r RepositoryMock) AddUnpaidToWorkedHours(employeeId models.ID) error {
	return getErrorByEmployeeId(employeeId)
}

func (r RepositoryMock) SetUnpaidHoursToZeroAndResetLastPayDate(employeeId models.ID) error {
	return getErrorByEmployeeId(employeeId)
}

func (r RepositoryMock) GetEmployeeConsumptions(employeeId models.ID) ([]models.Consumption, error) {
	err := getErrorByEmployeeId(employeeId)
	if err != nil {
		return nil, err
	}

	return getEmployeeIdToConsumptions()[employeeId], nil
}

func (r RepositoryMock) MakeConsumptionsCompleted(employeeId models.ID) error {
	return getErrorByEmployeeId(employeeId)
}

func getErrorByEmployeeId(employeeId models.ID) error {
	if employeeId == mock.BadEmployeeId {
		return mock.SomeError
	} else if employeeId == mock.NotFoundEmployeeId {
		return app_internals.EmployeeNotFound
	}

	return nil
}

func getEmployeeIdToConsumptions() map[models.ID][]models.Consumption {
	res := map[models.ID][]models.Consumption{}
	employees := mock.GetAllEmployees()
	consumptions := mock.GetAllConsumptions()

	for _, employee := range employees {
		var employeeConsumptions []models.Consumption
		for _, consumption := range consumptions {
			if consumption.EmployeeId == employee.Id {
				employeeConsumptions = append(employeeConsumptions, consumption)
			}
		}

		res[employee.Id] = employeeConsumptions
	}

	return res
}
