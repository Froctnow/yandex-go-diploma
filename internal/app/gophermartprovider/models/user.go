package models

type User struct {
	ID       string `db:"id"`
	Login    string `db:"login"`
	Password string `db:"password"`
}

type Withdraw struct {
	OrderNumber string  `db:"order_number"`
	Sum         float32 `db:"sum"`
	ProcessedAt string  `db:"processed_at"`
}
