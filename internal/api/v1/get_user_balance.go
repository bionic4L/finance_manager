package v1

import "github.com/gin-gonic/gin"

func getBalance(c *gin.Context) {
	// userID := c.Param("id")
	c.Status(200)
	c.Writer.Write([]byte("Это ручка 'getBalance' Я КРАСАВЧИК"))
}
