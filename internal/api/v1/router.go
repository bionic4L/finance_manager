package v1

import (
	"log"

	"github.com/gin-gonic/gin"
)

// TODO: get

func Router(address string) {
	router := gin.Default()

	router.GET("/balance", getBalance)
	router.GET("/confirmed/:user_id/:service_id/:order_id/:sum", paymentConfirmed)
	router.POST("/deposit", DepositToUser)
	router.POST("/reserve/:user_id/:service_id/:order_id/:price", reserveMoney)

	log.Print("running app...")
	router.Run(address)
}
