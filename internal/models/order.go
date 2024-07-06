package models

type Order struct {
	ID        int `json:"id"`
	ServiceID int `json:"service_id"`
	UserID    int `json:"user_id"`
	Sum       int `json:"sum"`
}
