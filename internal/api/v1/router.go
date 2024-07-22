package v1

import (
	"finance_manager/internal/repository"
	"finance_manager/internal/service"

	"github.com/gin-gonic/gin"
)

func Router(router *gin.Engine, repository *repository.Repository) {
	BalanceRouter(router, service.NewBalanceService(repository.BalanceRepository))
	DepositRouter(router, service.NewDepositService(repository.DepositRepository))
	CreateUserRouter(router, service.NewCreateUserService(repository.UserCreateRepository))
	router.DELETE("/user-delete", deleteUser)
}
