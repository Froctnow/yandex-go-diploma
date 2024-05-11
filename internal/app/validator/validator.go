package validator

type validator struct{}

type Validator interface {
}

func New() Validator {
	return &validator{}
}
