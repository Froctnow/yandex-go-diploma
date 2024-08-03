package models

type Order struct {
	Number     string   `json:"number"`
	Status     string   `json:"status"`
	Accrual    *float32 `json:"accrual"`
	UploadedAt string   `json:"uploaded_at"`
}
