package config

import "models"

const (
	hourlyRate  = "hourly"
	monthlyRate = "monthly"
)

func GetHourlyRateName() models.RateType {
	return hourlyRate
}

func GetMonthlyRateName() models.RateType {
	return monthlyRate
}
