package bookkeeping

import (
	"errors"
	"interfaces"
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

type BadPaymentFinalizerConstructor struct {
}

func (c BadPaymentFinalizerConstructor) GetPaymentFinalizer(
	models.PayType) interfaces.PaymentFinalizer {
	return BadPaymentFinalizer{}
}

func (c BadPaymentFinalizerConstructor) SetCheckProcessor(
	models.PayType, interfaces.CheckProcessor) {
	return
}

func (c BadPaymentFinalizerConstructor) SetDefaultCheckProcessor(interfaces.CheckProcessor) {
	return
}

type BadPaymentFinalizer struct {
}

func (f BadPaymentFinalizer) Finish(models.Payment) error {
	return errors.New("some-error")
}
