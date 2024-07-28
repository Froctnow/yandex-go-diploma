package models

type UseBalance struct {
	Balance   float32 `json:"balance"`
	Withdrawn float32 `json:"withdrawn"`
}

type UserWithdraw struct {
	OrderNumber string  `json:"order"`
	Sum         float32 `json:"sum"`
	ProcessedAt string  `json:"processed_at"`
}
