package dbactions

import (
	"errors"
	"log"
)

type Deposit interface {
	DepositToUser()
}

func (ud *User) DepositToUser(amount int) (User, error) {

	userTest := User{
		ID:      2,
		Balance: 4,
	}

	if userTest.Balance+amount < 0 {
		log.Printf("Отрицательный баланс: %d", userTest.Balance+amount)
		return userTest, errors.New("отрицательный баланс")
	}

	userTest.Balance = userTest.Balance + amount

	return userTest, nil
}
