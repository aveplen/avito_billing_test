package models

type BalanceRequest struct {
	UserID int64 `json:"user_id" binding:"required"`
}

type BalanceResponse struct {
	UserID int64  `json:"user_id"`
	Money  string `json:"money"`
}
