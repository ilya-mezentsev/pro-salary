package models

import (
	"time"
)

type Employee struct {
	Id          ID
	Name        Name
	PayType     PayType
	RateType    RateType
	Rate        Rate
	LastPayDate time.Time
}
