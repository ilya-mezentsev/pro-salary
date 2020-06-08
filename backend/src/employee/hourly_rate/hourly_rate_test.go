package hourly_rate

import (
	"models"
	utils "test_utils"
	"testing"
	"time"
)

var (
	testDate = time.Date(2020, 5, 30, 0, 0, 0, 0, time.Local)
)

func TestEmployee_IsPayDayAfterDay(t *testing.T) {
	e := New(models.Employee{LastPayDate: testDate})

	utils.AssertFalse(e.IsPayDay(utils.GetAfterOneDay(testDate)), t)
}

func TestEmployee_IsPayDayAfterWeek(t *testing.T) {
	e := New(models.Employee{LastPayDate: testDate})

	utils.AssertTrue(e.IsPayDay(utils.GetAfterOneWeek(testDate)), t)
}

func TestEmployee_IsPayDayAfterMonth(t *testing.T) {
	e := New(models.Employee{LastPayDate: testDate})

	utils.AssertTrue(e.IsPayDay(utils.GetAfterOneMonth(testDate)), t)
}

func TestEmployee_CalculatePayment40hours(t *testing.T) {
	employeeData := models.Employee{Rate: 500, UnpaidHours: 40}
	e := New(employeeData)

	utils.AssertEqual(
		models.Payment{
			EmployeeId: employeeData.Id,
			Amount:     models.PaymentAmount(employeeData.Rate * models.Rate(employeeData.UnpaidHours)),
		},
		e.CalculatePayment(),
		t,
	)
}

func TestEmployee_CalculatePayment30hours(t *testing.T) {
	employeeData := models.Employee{Rate: 500, UnpaidHours: 30}
	e := New(employeeData)

	utils.AssertEqual(
		models.Payment{
			EmployeeId: employeeData.Id,
			Amount:     models.PaymentAmount(employeeData.Rate * models.Rate(employeeData.UnpaidHours)),
		},
		e.CalculatePayment(),
		t,
	)
}

func TestEmployee_CalculatePayment50hours(t *testing.T) {
	employeeData := models.Employee{Rate: 500, UnpaidHours: 50}
	e := New(employeeData)

	utils.AssertEqual(
		models.Payment{
			EmployeeId: employeeData.Id,
			Amount:     models.PaymentAmount(employeeData.Rate * (40 + 10*1.5)),
		},
		e.CalculatePayment(),
		t,
	)
}

func TestEmployee_GetPaymentType(t *testing.T) {
	employeeData := models.Employee{PayType: "test"}
	e := New(employeeData)

	utils.AssertEqual(employeeData.PayType, e.GetPaymentType(), t)
}
