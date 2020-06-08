package finalizer

import (
	"mock"
	. "mock/payment_finalizer"
	"models"
	"payment_finalizer"
	utils "test_utils"
	"testing"
)

var (
	defaultCheckProcessor      = payment_finalizer.DefaultCheckProcessor{}
	paymentFinalizerRepository = RepositoryMock{}
	finalizer                  = New(
		paymentFinalizerRepository,
		defaultCheckProcessor,
	)
)

func TestFinalizer_FinishSuccess(t *testing.T) {
	err := finalizer.Finish(models.Payment{
		EmployeeId: mock.GetAllEmployees()[0].Id,
		Amount:     100,
	})

	utils.AssertNil(err, t)
}

func TestFinalizer_FinishSomeError(t *testing.T) {
	err := finalizer.Finish(models.Payment{
		EmployeeId: mock.BadEmployeeId,
		Amount:     100,
	})

	utils.AssertNotNil(err, t)
}
