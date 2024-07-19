package dbactions

type Balance interface {
	GetUserBalance()
}

func (ud *User) GetUserBalance() (User, error) {
	return User{
		ID:      1,
		Balance: 777,
	}, nil
}
