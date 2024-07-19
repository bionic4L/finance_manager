package models

type Deposit struct {
	DepositID     int `json:"deposit_id"`
	UserID        int `json:"user_id"`
	DepositAmount int `json:"amount"`
}
