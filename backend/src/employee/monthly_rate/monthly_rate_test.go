package monthly_rate

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

	utils.AssertFalse(e.IsPayDay(utils.GetAfterOneWeek(testDate)), t)
}

func TestEmployee_IsPayDayAfterMonth(t *testing.T) {
	e := New(models.Employee{LastPayDate: testDate})

	utils.AssertTrue(e.IsPayDay(utils.GetAfterOneMonth(testDate)), t)
}
