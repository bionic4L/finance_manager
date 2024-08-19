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
	defer c.Request.Body.Close()
	if err != nil {
		c.Status(400)
		log.Error("ошибка чтения тела запроса: ", err)
		return
	}

	if err := json.Unmarshal(jsonRequestBody, &dep); err != nil {
		c.JSON(400, "проверьте правильность введенных данных")
		log.Error("ошибка декодирования json: ", err)
		return
	}

	if err := ValidateDeposit(c, dep); err != nil {
		c.JSON(400, "валидация запроса не пройдена")
		log.Error("ошибка валидации: ", err)
		return
	}

	if err := d.service.Deposit(ctx, dep.UserID, dep.DepositAmount); err != nil {
		log.Error("ошибка при выполнении депозита: ", err)
		c.JSON(400, "ошибка при выполнении депозита")
		return
	}
	c.JSON(200, "успешный депозит")

}

func ValidateDeposit(c *gin.Context, d *models.Deposit) error {
	if c.Request.Header.Get("Content-Type") != "application/json" {
		c.Status(415)
		return errors.New("неподдерживаемый тип контента")
	}

	if d.DepositAmount < 0 {
		return errors.New("депозит не может быть отрицательным числом")
	}

	if d.UserID <= 0 {
		return errors.New("id пользователя не может быть отрицательным числом")
	}

	//...

	return nil
}
