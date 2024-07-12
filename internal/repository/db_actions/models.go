package dbactions

type (
	User struct {
		ID      int `json:"id"`
		Balance int `json:"balance"`
	}

	TransactionInfo struct {
		ID     int `json:"id"`
		FromID int `json:"from_id"`
		ToID   int `json:"to_id"`
		Amount int `json:"amount"`
	}
)
