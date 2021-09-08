package models

type Withdrawal struct {
	UserID int64  `json:"user_id" binding:"required"`
	Money  string `json:"money" binding:"required"`
}
