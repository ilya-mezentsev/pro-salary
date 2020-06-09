package bookkeeping

import (
	"models"
)

type SaveChecksProcessor struct {
	checks []models.Check
}

func (p *SaveChecksProcessor) Process(check models.Check) error {
	p.checks = append(p.checks, check)
	return nil
}

func (p SaveChecksProcessor) GetSavedChecks() []models.Check {
	return p.checks
}
