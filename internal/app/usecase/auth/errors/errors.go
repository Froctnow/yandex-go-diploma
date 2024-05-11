package errors

import "fmt"

type UserAlreadyExistsError struct {
}

func (e UserAlreadyExistsError) Error() string {
	return fmt.Sprintf("user already exists")
}
