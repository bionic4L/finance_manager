package v1

import (
	"encoding/json"
	"errors"
	"finance_manager/internal/models"
	dbactions "finance_manager/internal/repository/db_actions"
	"finance_manager/internal/service"
	"io"
	"log"

	"github.com/gin-gonic/gin"
)

type Deposit struct {
	service *service.DepositService
}

func DepositRouter(r *gin.Engine, service *service.DepositService) {
	d := &Deposit{service: service}

	r.POST("/deposit", d.DepositToUser)
}

func (d Deposit) DepositToUser(c *gin.Context) {
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

	jsonRequestBody, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.Status(400)
		return errors.New("ошибка чтения тела запроса")
	}

	if err := json.Unmarshal(jsonRequestBody, &dep); err != nil {
		c.Status(400)
		return errors.New("ошибка декодирования json")
	}

	if c.Request.Header.Get("Content-Type") != "application/json" {
		c.Status(415)
		return errors.New("неподдерживаемый тип контента")
	}

	db := dbactions.DepositRepository{}
	if err := db.Deposit(dep.UserID, dep.DepositAmount); err != nil {
		return err
	}

	// if err != nil {
	// 	c.Status(422)
	// 	c.Writer.Write([]byte("ошибка: средства не были зачислены"))
	// 	return errors.New("ошибка: средства не были зачислены")
	// }

	c.JSON(200, dep)

	return nil
}
