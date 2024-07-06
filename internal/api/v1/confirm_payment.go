package v1

import "github.com/gin-gonic/gin"

func paymentConfirmed(c *gin.Context) {
	userID := c.Param("user_id")
	serviceID := c.Param("service_id")
	orderID := c.Param("order_id")
	sum := c.Param("sum")
}
