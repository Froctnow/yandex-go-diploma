package accrualservice

type GetOrderResponse struct {
	// Номер заказа
	Order   uint32  `json:"order"`
	Status  string  `json:"status"`
	Accrual *uint32 `json:"accrual,omitempty"`
}
