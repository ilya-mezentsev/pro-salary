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

func New(repository interfaces.EmployeesRepository) interfaces.EmployeesFactory {
	return Factory{repository}
}

func (f Factory) GetAll() ([]interfaces.Employee, error) {
	var employees []interfaces.Employee
	employeesData, err := f.repository.GetAll()
	if err != nil {
		return nil, err
	}

	for _, employeeData := range employeesData {
		if f.isHourlyRate(employeeData) {
			employees = append(employees, hourly_rate.New(employeeData))
		} else if f.isMonthlyRate(employeeData) {
			employees = append(employees, monthly_rate.New(employeeData))
		}
	}

	return employees, nil
}

func (f Factory) isHourlyRate(data models.Employee) bool {
	return data.RateType == config.GetHourlyRateName()
}

func (f Factory) isMonthlyRate(data models.Employee) bool {
	return data.RateType == config.GetMonthlyRateName()
}
