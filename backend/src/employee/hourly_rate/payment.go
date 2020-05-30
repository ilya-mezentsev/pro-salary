package hourly_rate

import "models"

const extraHoursCoefficient = 1.5

type payment struct {
	data models.Employee
}

func (p payment) calculate(workedHours int) models.Payment {
	if workedHours > weekWorkHours {
		return models.Payment(
			p.data.Rate*weekWorkHours +
				p.data.Rate*models.Rate(extraHoursCoefficient)*models.Rate(workedHours%weekWorkHours))
	} else {
		return models.Payment(p.data.Rate * models.Rate(workedHours))
	}
}
