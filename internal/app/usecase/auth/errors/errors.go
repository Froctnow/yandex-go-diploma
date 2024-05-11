package errors

type UserAlreadyExistsError struct {
}

func (e UserAlreadyExistsError) Error() string {
	return "user already exists"
}

type IncorrectLoginPasswordError struct {
}

func (e IncorrectLoginPasswordError) Error() string {
	return "incorrect login/password"
}
