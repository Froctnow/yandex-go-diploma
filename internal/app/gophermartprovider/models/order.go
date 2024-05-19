package models

type Order struct {
	Number     string  `db:"number"`
	Status     string  `db:"status"`
	Accrual    *uint32 `db:"accrual"`
	UploadedAt string  `db:"uploaded_at"`
}
