package v1

import (
	"finance_manager/internal/service"
	"log"

	"github.com/gin-gonic/gin"
)

func Router(address string) {
	router := gin.Default()

	BalanceRouter(router, service.BalanceService{})
	// router.POST("/transaction", Transaction)
	router.POST("/deposit", DepositToUser)
	router.POST("/user-add", addUser)
	router.DELETE("/user-delete", deleteUser)

	log.Print("running app...")
	router.Run(address)
}
