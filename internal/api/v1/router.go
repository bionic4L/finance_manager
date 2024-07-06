package v1

import "github.com/gin-gonic/gin"

func Router() {
	router := gin.Default()

	router.GET("/balance/:id", getBalance)
	router.POST("/deposit/:id", addMoney)
	router.POST("/reserve/:user_id/:service_id/:order_id/:price", reserveMoney)
	router.GET("/confirmed/:user_id/:service_id/:order_id/:sum", paymentConfirmed)

	router.Run("localhost:8080")
}
