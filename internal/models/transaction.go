package models

type Transaction struct {
	ID     int `json:"id"`
	FromID int `json:"from_id"`
	ToID   int `json:"to_id"`
	Amount int `json:"amount"`
}
