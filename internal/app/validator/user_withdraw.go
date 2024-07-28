package validator

import (
	"strconv"

	"github.com/gobuffalo/validate"
)

func (v *validator) UserWithdraw(orderNumber string) *validate.Errors {
	numericOrderNumber, err := strconv.Atoi(orderNumber)
	if err != nil {
		errors := validate.NewErrors()

		errors.Add("OrderNumber", "Order number must be a number")

		return errors
	}

	checks := []validate.Validator{
		&AlgorithmLuhn{
			Name:   "OrderNumber",
			Number: numericOrderNumber,
		},
	}
	errors := validate.Validate(checks...)
	return errors
}
