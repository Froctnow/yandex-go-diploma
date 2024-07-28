package models

type RegisterRequest struct {
	Login    string `json:"login" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginRequest struct {
	Login    string `json:"login" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type GetUserBalanceResponse struct {
	Balance   float32 `json:"balance"`
	Withdrawn float32 `json:"withdrawn"`
}

type WithdrawRequest struct {
	Sum   float32 `json:"sum"`
	Order string  `json:"order"`
}
