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

func DepositToUser(c *gin.Context) {
	if err := ValidateDeposit(c); err != nil {
		c.Status(400)
		log.Print("валидация запроса не пройдена")
		log.Print(err)
		return
	}

	c.Status(200)

}

func ValidateDeposit(c *gin.Context) error {
	var dep models.Deposit

	JSONRequestBody, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.Status(400)
		return errors.New("ошибка чтения тела запроса")
	}

	if err := json.Unmarshal(JSONRequestBody, &dep); err != nil {
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

	db := dbactions.User{}
	ud, err := db.DepositToUser(dep.DepositAmount)
	if err != nil {
		c.Status(422)
		c.Writer.Write([]byte("ошибка: средства не были зачислены"))
		return errors.New("ошибка: средства не были зачислены")
	}
	if ud.ID != dep.UserID {
		c.Status(404)
		c.Writer.Write([]byte("пользователь с таким id не найден"))
		return errors.New("пользователь с таким id не найден")
	}

	c.JSON(200, ud)

	return nil
}
