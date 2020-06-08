package hourly_rate

import "models"

const extraHoursCoefficient = 1.5

type payment struct {
	data models.Employee
}

func (p payment) calculate(workedHours int) models.Payment {
	payment := models.Payment{
		EmployeeId: p.data.Id,
	}

	if workedHours > weekWorkHours {
		payment.Amount = models.PaymentAmount(p.data.Rate*weekWorkHours +
			p.data.Rate*models.Rate(extraHoursCoefficient)*models.Rate(workedHours%weekWorkHours))
	} else {
		payment.Amount = models.PaymentAmount(p.data.Rate * models.Rate(workedHours))
	}

	return payment
}
