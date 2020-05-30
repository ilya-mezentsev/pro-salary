package models

import (
	"time"
)

type Employee struct {
	Id          ID        `db:"uuid"`
	Name        Name      `db:"name"`
	PayType     PayType   `db:"pay_type"`
	RateType    RateType  `db:"rate_type"`
	Rate        Rate      `db:"rate"`
	LastPayDate time.Time `db:"last_pay_date"`
}
