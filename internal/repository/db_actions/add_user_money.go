package dbactions

import (
	"errors"
	"log"
)

func (ud *User) DepositToUser(amount int) (User, error) {

	userTest := User{
		ID:      2,
		Balance: 4,
	}

	finalBalance := userTest.Balance + amount
	if finalBalance < 0 {
		log.Printf("Отрицательный баланс: %d", finalBalance)
		return userTest, errors.New("отрицательный баланс")
	}

	userTest.Balance = finalBalance

	return userTest, nil
}
