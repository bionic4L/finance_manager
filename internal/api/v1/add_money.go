package v1

import "github.com/gin-gonic/gin"

func addMoney(c *gin.Context) {
	// userID := c.Param("id")
	c.Status(200)
	c.Writer.Write([]byte("Это ручка 'addMoney' Я КРАСАВЧИК"))
}
