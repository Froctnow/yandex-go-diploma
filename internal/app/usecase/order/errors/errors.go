package errors

type UserOrderAlreadyExists struct{}

func (e UserOrderAlreadyExists) Error() string {
	return "order of user already exists"
}

type OrderAlreadyExists struct{}

func (e OrderAlreadyExists) Error() string {
	return "order already exists"
}
