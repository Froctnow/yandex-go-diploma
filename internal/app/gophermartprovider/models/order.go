package models

type Order struct {
	Number     string  `db:"number"`
	Status     string  `db:"status"`
	Accrual    *uint32 `db:"accrual,omitempty"`
	UploadedAt string  `db:"uploaded_at"`
}
