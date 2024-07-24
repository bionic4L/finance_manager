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

	r.PATCH("/transaction", t.Transaction)
}

func (t *Transaction) Transaction(c *gin.Context) {
	var transactionModel *models.Transaction

	JSONRequestBody, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.Status(400)
		log.Error(errors.New("ошибка чтения тела запроса"))
		return
	}

	if err := json.Unmarshal(JSONRequestBody, &transactionModel); err != nil {
		c.Status(400)
		log.Error(errors.New("ошибка декодирования json"))
		return
	}

	if err := t.service.Transaction(transactionModel.FromID, transactionModel.ToID, transactionModel.Amount); err != nil {
		c.Status(400)
		c.Writer.Write([]byte("ошибка во время выполнения транзакции"))
		log.Error(err)
		return
	}

	if err := ValidateTransaction(c, transactionModel); err != nil {
		c.Status(400)
		log.Warn("валидация запроса не пройдена")
		return
	}

	c.Writer.Write([]byte("перевод выполнен!"))
	c.Status(200)
}

func ValidateTransaction(c *gin.Context, transactionModel *models.Transaction) error {
	return nil
}
