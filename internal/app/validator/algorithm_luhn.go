package validator

import (
	"strings"

	"github.com/gobuffalo/validate"
)

type AlgorithmLuhn struct {
	Name    string
	Number  int
	Message string
}

func (v *AlgorithmLuhn) IsValid(errors *validate.Errors) {
	if v.Message == "" {
		v.Message = "Number is not valid by Luhn algorithm"
	}

	if (v.Number%10+checksum(v.Number/10))%10 != 0 {
		errors.Add(strings.ToLower(v.Name), v.Message)
	}
}

func checksum(number int) int {
	var luhn int

	for i := 0; number > 0; i++ {
		cur := number % 10

		if i%2 == 0 { // even
			cur = cur * 2
			if cur > 9 {
				cur = cur%10 + cur/10
			}
		}

		luhn += cur
		number = number / 10
	}
	return luhn % 10
}
