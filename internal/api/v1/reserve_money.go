package v1

import "github.com/gin-gonic/gin"

func reserveMoney(c *gin.Context) {
	// userID := c.Param("user_id")
	// serviceID := c.Param("service_id")
	// orderID := c.Param("order_id")
	// price := c.Param("price")
	c.Status(200)
	c.Writer.Write([]byte("Это ручка 'reserveMoney' Я КРАСАВЧИК"))
}
