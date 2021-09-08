package models

type Transaction struct {
	FromID int64  `json:"from_id" binding:"required"`
	ToID   int64  `json:"to_id" binding:"required"`
	Money  string `json:"money" binding:"required"`
}
