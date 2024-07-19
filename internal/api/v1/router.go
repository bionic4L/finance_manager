package v1

import (
	"log"

	"github.com/gin-gonic/gin"
)

func Router(address string) {
	router := gin.Default()

	router.GET("/balance", getBalance)
	router.POST("/transaction", Transaction)
	router.POST("/deposit", DepositToUser)
	router.POST("/reserve", reserveMoney)

	log.Print("running app...")
	router.Run(address)
}
