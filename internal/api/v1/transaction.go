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

type Transaction struct {
	service *service.TransactionService
}

func TransactionRouter(r *gin.Engine, service *service.TransactionService) {
	t := Transaction{service: service}

	r.POST("/transaction", t.Transaction)
}

func (t *Transaction) Transaction(c *gin.Context) {
	var transactionModel *models.Transaction
	ctx := c.Request.Context()

	JSONRequestBody, err := io.ReadAll(c.Request.Body)
	defer c.Request.Body.Close()
	if err != nil {
		c.Status(400)
		log.Error("ошибка чтения тела запроса", err)
		return
	}

	if err := json.Unmarshal(JSONRequestBody, &transactionModel); err != nil {
		c.JSON(400, "проверьте правильность введенных данных")
		log.Error("ошибка декодирования json", err)
		return
	}

	if err := ValidateTransaction(c, transactionModel); err != nil {
		c.JSON(400, "валидация запроса не пройдена")
		log.Error("ошибка валидации: ", err)
		return
	}

	if err := t.service.Transaction(ctx, transactionModel.FromID, transactionModel.ToID, transactionModel.Amount); err != nil {
		c.JSON(400, "ошибка во время выполнения транзакции")
		log.Error("ошибка во время выполнения транзакции", err)
		return
	}
	c.JSON(200, "перевод выполнен!")
}

func ValidateTransaction(c *gin.Context, transactionModel *models.Transaction) error {
	switch {
	case transactionModel.Amount < 1:
		return errors.New("сумма транзакции не может быть меньше 1")
	case transactionModel.FromID < 1 || transactionModel.ToID < 1:
		return errors.New("id пользователя не может быть меньше 1")
	default:
		return nil
	}
}
