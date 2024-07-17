package dbactions

import (
	"errors"
	"log"
)

type Transaction interface {
	Transaction()
}

func (t *TransactionInfo) Transaction(from int, to int, amount int) (TransactionInfo, User, error) {
	user1 := User{
		ID:      1,
		Balance: 1000,
	}

	user2 := User{
		ID:      2,
		Balance: 0,
	}

	newTransaction := TransactionInfo{
		ID:     1,
		FromID: from,
		ToID:   to,
		Amount: amount,
	}

	if from == to {
		log.Print("нельзя перевести средства самому себе")
		return newTransaction, user2, errors.New("нельзя перевести средства самому себе")
	}

	if from != user1.ID && from != user2.ID || to != user1.ID && to != user2.ID {
		log.Print("пользователь с таким id не найден")
		return newTransaction, user2, errors.New("пользователь с таким id не найден")
	}

	if user1.ID == from {
		if user1.Balance < user2.Balance {
			log.Printf("отрицательный баланс: %d", user1.Balance-amount)
			return newTransaction, user2, errors.New("отрицательный баланс")
		}
		user1.Balance -= amount
		user2.Balance += amount
		return newTransaction, user2, nil
	}

	if user2.ID == from {
		if user2.Balance < user1.Balance {
			log.Printf("Отрицательный баланс: %d", user2.Balance-amount)
			return newTransaction, user2, errors.New("отрицательный баланс")
		}
		user1.Balance += amount
		user2.Balance -= amount
		return newTransaction, user2, nil
	}

	return newTransaction, user2, nil
}
