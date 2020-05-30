package employee

import (
	"config"
	"employee/hourly_rate"
	"employee/monthly_rate"
	"interfaces"
	"models"
)

type Factory struct {
	repository interfaces.EmployeesRepository
}

func New(repository interfaces.EmployeesRepository) Factory {
	return Factory{repository}
}

func (f Factory) GetAll() ([]interfaces.Employee, error) {
	var employees []interfaces.Employee
	employeesData, err := f.repository.GetAll()
	if err != nil {
		return nil, err
	}

	for _, employeeData := range employeesData {
		if f.isEmployeeWithHourlyRate(employeeData) {
			employees = append(employees, hourly_rate.New(employeeData))
		} else if f.isEmployeeWithMonthlyRate(employeeData) {
			employees = append(employees, monthly_rate.New(employeeData))
		}
	}

	return employees, nil
}

func (f Factory) isEmployeeWithHourlyRate(data models.Employee) bool {
	return data.RateType == config.GetHourlyRateName()
}

func (f Factory) isEmployeeWithMonthlyRate(data models.Employee) bool {
	return data.RateType == config.GetMonthlyRateName()
}
