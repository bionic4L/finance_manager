package models

type StatusController struct {
	TransactionID int  `json:"transaction_id"`
	Confirm       bool `json:"confirm"`
}
