package models

type Order struct {
	Id          int `json:"id"`
	Service_ID  int `json:"service_id"`
	User_ID     int `json:"user_id"`
	Price_Total int `json:"price_total"`
}
