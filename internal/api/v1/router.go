package v1

import (
	"finance_manager/internal/repository"
	"finance_manager/internal/service"

	"github.com/gin-gonic/gin"
)

func Router(router *gin.Engine, repository *repository.Repository) {
	BalanceRouter(router, service.NewBalanceService(repository.BalanceRepository))
	// router.POST("/transaction", Transaction)
	router.POST("/deposit", DepositToUser)
	router.POST("/user-add", addUser)
	router.DELETE("/user-delete", deleteUser)
}
