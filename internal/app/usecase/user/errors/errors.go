package errors

type UserNotEnoughBalance struct{}

func (e UserNotEnoughBalance) Error() string {
	return "user balance is not enough"
}
