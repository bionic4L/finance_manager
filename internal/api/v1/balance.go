package v1

import (
	"errors"
	"finance_manager/internal/service"
	"strconv"

	log "github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
)

type Balance struct {
	service *service.BalanceService
}

func BalanceRouter(r *gin.Engine, service *service.BalanceService) {
	b := &Balance{service: service}

	r.GET("/balance", b.getBalance)
}

func (b *Balance) getBalance(c *gin.Context) {
	ctx := c.Request.Context()

	err := ValidateGetBalance(c)
	if err != nil {
		return
	}
	userID, _ := strconv.Atoi(c.Query("id"))

	userData, err := b.service.GetBalance(ctx, userID) //прокид с транспортного уровня на сервисный
	if err != nil {
		log.Print(err)
		return
	}

	if userData.ID != userID {
		c.Status(404)
		c.Writer.Write([]byte("пользователь с таким id не найден"))
		return
	}

	c.Status(200)
	c.JSON(200, userData)
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
