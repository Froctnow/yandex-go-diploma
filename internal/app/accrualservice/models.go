package accrualservice

type GetOrderResponse struct {
	// Номер заказа
	Order   uint32   `json:"order"`
	Status  string   `json:"status"`
	Accrual *float32 `json:"accrual,omitempty"`
}
