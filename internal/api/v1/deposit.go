package v1

import (
	"encoding/json"
	"errors"
	"finance_manager/internal/models"
	"finance_manager/internal/service"
	"io"

	log "github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
)

type Deposit struct {
	service *service.DepositService
}

func DepositRouter(r *gin.Engine, service *service.DepositService) {
	d := &Deposit{service: service}

	r.POST("/deposit", d.Deposit)
}

func (d *Deposit) Deposit(c *gin.Context) {
	var dep *models.Deposit
	ctx := c.Request.Context()

	jsonRequestBody, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.Status(400)
		log.Error(errors.New("ошибка чтения тела запроса"))
		return
	}

	if err := json.Unmarshal(jsonRequestBody, &dep); err != nil {
		c.Status(400)
		c.Writer.Write([]byte("убедитесь, что вы ввели корректные данные"))
		log.Print(errors.New("ошибка декодирования json"))
		return
	}

	if err := ValidateDeposit(c, dep); err != nil {
		c.Status(400)
		log.Warn("валидация запроса не пройдена")
		log.Error(err)
		return
	}

	if err := d.service.Deposit(ctx, dep.UserID, dep.DepositAmount); err != nil {
		log.Error(err)
		c.Status(400)
		return
	}

	c.Status(200)
	// c.JSON(200, dep)
	c.Writer.Write([]byte("успешный депозит"))

}

func ValidateDeposit(c *gin.Context, d *models.Deposit) error {
	if c.Request.Header.Get("Content-Type") != "application/json" {
		c.Status(415)
		return errors.New("неподдерживаемый тип контента")
	}

	if d.DepositAmount < 0 {
		c.Status(400)
		c.Writer.Write([]byte("депозит не может быть отрицательным числом"))
		return errors.New("депозит не может быть отрицательным числом")
	}

	if d.UserID <= 0 {
		c.Status(400)
		c.Writer.Write([]byte("id пользователя не может быть отрицательным числом"))
		return errors.New("id пользователя не может быть отрицательным числом")
	}

	//...

	return nil
}
