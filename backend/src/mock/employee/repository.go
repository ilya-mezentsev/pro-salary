package employee

import (
	"mock"
	"models"
)

type RepositoryMock struct {
	Error error
}

func (r RepositoryMock) GetAll() ([]models.Employee, error) {
	if r.Error != nil {
		return nil, r.Error
	}

	return mock.GetAllEmployees(), nil
}
