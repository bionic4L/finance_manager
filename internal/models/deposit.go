package models

type Deposit struct {
	DepositID     int `json:"deposit_id" db:"deposit_id"`
	UserID        int `json:"user_id" db:"user_id"`
	DepositAmount int `json:"amount" db:"amount"`
}
