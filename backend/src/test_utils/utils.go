package utils

import (
	"fmt"
	"os"
	"testing"
	"time"
)

type Expectation struct {
	Expected, Actual interface{}
}

func SkipInShortMode() {
	if os.Getenv("SHORT_MODE") == "1" {
		fmt.Println("skipping DB tests in short mode")
		os.Exit(0)
	}
}

func AssertTrue(actual bool, t *testing.T) {
	if !actual {
		logExpectationAndFail(true, actual, t)
	}
}

func AssertFalse(actual bool, t *testing.T) {
	if actual {
		logExpectationAndFail(false, actual, t)
	}
}

func AssertEqual(expected, actual interface{}, t *testing.T) {
	if expected != actual {
		logExpectationAndFail(expected, actual, t)
	}
}

func AssertNotEqual(notExpected, actual interface{}, t *testing.T) {
	if notExpected == actual {
		logExpectationAndFail(fmt.Sprintf("not equal to %v", notExpected), actual, t)
	}
}

func AssertNil(v interface{}, t *testing.T) {
	if v != nil {
		logExpectationAndFail(nil, v, t)
	}
}

func AssertNotNil(v interface{}, t *testing.T) {
	if v == nil {
		logExpectationAndFail("not nil", v, t)
	}
}

func AssertErrorsEqual(expectedErr, actualErr error, t *testing.T) {
	if expectedErr != actualErr {
		logExpectationAndFail(expectedErr, actualErr, t)
	}
}

func logExpectationAndFail(expected, actual interface{}, t *testing.T) {
	t.Log(
		GetExpectationString(
			Expectation{Expected: expected, Actual: actual}))
	t.Fail()
}

func GetExpectationString(e Expectation) string {
	return fmt.Sprintf("expected: %v, got: %v\n", e.Expected, e.Actual)
}

func GetAfterOneDay(date time.Time) time.Time {
	return date.AddDate(0, 0, 1)
}

func GetAfterOneWeek(date time.Time) time.Time {
	return date.AddDate(0, 0, 7)
}

func GetAfterOneMonth(date time.Time) time.Time {
	return date.AddDate(0, 1, 0)
}
