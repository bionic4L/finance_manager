package models

type Report struct {
	ID          int    `json:"id"`
	ServiceName string `json:"service_name"`
	Profit      int    `json:"profit"`
}
