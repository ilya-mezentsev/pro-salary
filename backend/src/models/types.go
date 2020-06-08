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
		EmployeeId ID
		Amount     PaymentAmount
	}

	Consumption struct {
		Id         ID `db:"uuid"`
		EmployeeId ID `db:"employee_uuid"`
		Amount     PaymentAmount
	}

	Check struct {
		EmployeeId   ID
		Amount       PaymentAmount
		Consumptions []Consumption
		Total        PaymentAmount
	}
)
