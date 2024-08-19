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

	if err := ValidateGetBalance(c); err != nil {
		c.JSON(400, "валидация запроса не пройдена")
		log.Error("ошибка валидации запроса: ", err)
		return
	}
	userID, _ := strconv.Atoi(c.Query("id"))

	userData, err := b.service.GetBalance(ctx, userID) //прокид с транспортного уровня на сервисный
	if err != nil {
		c.JSON(400, "ошибка во время получения баланса")
		log.Error(err)
		return
	}

	if userData.ID != userID {
		c.JSON(404, "пользователь с таким id не найден")
		return
	}
	c.JSON(200, userData)
}

func ValidateGetBalance(c *gin.Context) error {
	userID := c.Query("id")

	if userID == "" {
		return errors.New("вы забыли указать параметр 'id'")
	}

	idFig, err := strconv.Atoi(userID)

	if err != nil {
		return errors.New("параметр 'id' должен быть цифрой")
	}

	if idFig < 0 {
		return errors.New("id не может быть отрицательным")
	}

	return nil
}
