package models

type User struct {
	ID      int    `json:"id" db:"id"`
	Name    string `json:"name" db:"name"`
	Balance int    `json:"balance" db:"balance"`
}
