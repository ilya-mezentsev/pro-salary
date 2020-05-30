package types

type (
	PaymentProcessing struct {
		Error chan error
		Done  chan bool
	}
)
