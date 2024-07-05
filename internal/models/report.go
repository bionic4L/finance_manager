package models

type Report struct {
	Id           int    `json:"id"`
	Service_name string `json:"service_name"`
	Profit       int    `json:"profit"`
}
