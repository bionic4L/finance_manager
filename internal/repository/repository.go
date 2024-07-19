package repository

type Repository interface {
	GetUserBalance()
	DepositToUser()
	Transaction()
}
