package models

type Deposit struct {
	ID            int `json:"id"`
	UserID        int `json:"user_id"`
	DepositAmount int `json:"amount"`
}
