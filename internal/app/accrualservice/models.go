package accrualservice

type GetOrderResponse struct {
	// Номер заказа
	Order   string   `json:"order"`
	Status  string   `json:"status"`
	Accrual *float32 `json:"accrual,omitempty"`
}
