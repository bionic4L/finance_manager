package v1

import (
	"errors"
	dbactions "finance_manager/internal/repository/db_actions"
	"strconv"

	"github.com/gin-gonic/gin"
)

// type Balance struct {
// 	service service.BalanceService
// }

// func BalanceRouter(r *gin.Engine, s service.BalanceService) {
// 	b := &Balance{service: s}

// 	r.GET("balance/:id", b.getBalance())
// }

func getBalance(c *gin.Context) {
	err := ValidateGetBalance(c)
	if err != nil {
		return
	}
	userID := c.Query("id")
	IDInt, _ := strconv.Atoi(userID)

	db := dbactions.User{}
	ud, _ := db.GetUserBalance()

	if ud.ID != IDInt {
		c.Status(404)
		c.Writer.Write([]byte("пользователь с таким id не найден"))
		return
	}

	c.Status(200)
	c.JSON(200, ud)
}

func ValidateGetBalance(c *gin.Context) error {
	userID := c.Query("id")

	if userID == "" {
		c.Status(422)
		c.Writer.Write([]byte("вы забыли указать параметр 'id'"))
		return errors.New("вы забыли указать параметр 'id'")
	}

	idFig, err := strconv.Atoi(userID)

	if err != nil {
		c.Status(422)
		c.Writer.Write([]byte("параметр 'id' должен быть цифрой."))
		return errors.New("параметр 'id' должен быть цифрой")
	}

	if idFig < 0 {
		c.Status(422)
		c.Writer.Write([]byte("id не может быть отрицательным"))
		return errors.New("id не может быть отрицательным")
	}

	return nil
}
