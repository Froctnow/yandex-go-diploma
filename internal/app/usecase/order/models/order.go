package models

type Order struct {
	Number     string  `json:"number"`
	Status     string  `json:"status"`
	Accrual    *uint32 `json:"accrual"`
	UploadedAt string  `json:"uploaded_at"`
}
