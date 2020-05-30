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
	employeeData := models.Employee{Rate: 500}
	e := New(employeeData)
	workedHours := 40

	utils.AssertEqual(
		models.Payment(employeeData.Rate*models.Rate(workedHours)),
		e.CalculatePayment(workedHours),
		t,
	)
}

func TestEmployee_CalculatePayment30hours(t *testing.T) {
	employeeData := models.Employee{Rate: 500}
	e := New(employeeData)
	workedHours := 30

	utils.AssertEqual(
		models.Payment(employeeData.Rate*models.Rate(workedHours)),
		e.CalculatePayment(workedHours),
		t,
	)
}

func TestEmployee_CalculatePayment50hours(t *testing.T) {
	employeeData := models.Employee{Rate: 500}
	e := New(employeeData)
	workedHours := 50

	utils.AssertEqual(
		models.Payment(employeeData.Rate*40)+models.Payment(employeeData.Rate*10*1.5),
		e.CalculatePayment(workedHours),
		t,
	)
}
