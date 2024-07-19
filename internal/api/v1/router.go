package v1

import (
	"finance_manager/internal/repository"
	"finance_manager/internal/service"

	"github.com/gin-gonic/gin"
)

func Router(router *gin.Engine, repository *repository.Repository) {
	BalanceRouter(router, service.NewBalanceService(repository.BalanceRepository))
	DepositRouter(router, service.NewDepositService(repository.DepositRepository))
	// router.POST("/transaction", Transaction)
	router.POST("/user-add", addUser)
	router.DELETE("/user-delete", deleteUser)
}
