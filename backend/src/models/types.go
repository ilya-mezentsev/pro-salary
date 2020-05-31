package models

type (
	ID            string
	Name          string
	PayType       string
	RateType      string
	Rate          float64
	PaymentAmount float64
)

type (
	Payment struct {
		UserId ID
		Amount PaymentAmount
	}
)
