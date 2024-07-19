package v1

import (
	"encoding/json"
	"errors"
	"finance_manager/internal/models"
	dbactions "finance_manager/internal/repository/db_actions"
	"io"
	"log"

	"github.com/gin-gonic/gin"
)

func Transaction(c *gin.Context) {
	if err := ValidateTransaction(c); err != nil {
		c.Status(400)
		log.Print("валидация запроса не пройдена")
		return
	}
	c.Status(200)
}

func ValidateTransaction(c *gin.Context) error {
	var t models.Transaction

	JSONRequestBody, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.Status(400)
		return errors.New("ошибка чтения тела запроса")
	}

	if err := json.Unmarshal(JSONRequestBody, &t); err != nil {
		c.Status(400)
		return errors.New("ошибка декодирования json")
	}

	if c.Request.Method != "POST" {
		c.Status(405)
		return errors.New("метод не поддерживается")
	}

	if c.Request.Header.Get("Content-Type") != "application/json" {
		c.Status(415)
		return errors.New("неподдерживаемый тип контента")
	}

	db := dbactions.TransactionInfo{}
	transaction, user, err := db.Transaction(t.FromID, t.ToID, t.Amount)
	if t.FromID != transaction.FromID || t.ToID != transaction.ToID {
		c.Status(400)
		c.Writer.Write([]byte("пользователь(и) с таким id не найдены"))
		log.Print(t.FromID)
		log.Print(t.ToID)
		log.Print(transaction.FromID)
		log.Print(transaction.ToID)
		return errors.New("пользователь(и) с таким id не найдены")
	}

	if t.FromID == t.ToID {
		c.Status(400)
		c.Writer.Write([]byte("нельзя перевести средства самому себе"))
		return errors.New("нельзя перевести средства самому себе")
	}

	if err != nil {
		c.Status(400)
		c.Writer.Write([]byte("ошибка во время выполнения транзакции"))
		return errors.New("ошибка во время выполнения транзакции")
	}

	c.JSON(200, transaction)
	c.JSON(200, user)

	return nil
}
