package v1

import (
	"encoding/json"
	"errors"
	"finance_manager/internal/models"
	"finance_manager/internal/service"
	"io"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type StatusController struct {
	service *service.StatusControllerService
}

func StatusControllerRouter(r *gin.Engine, service *service.StatusControllerService) {
	sc := &StatusController{service: service}

	r.POST("/status", sc.StatusController)
}

func (sc *StatusController) StatusController(c *gin.Context) {
	var s *models.StatusController
	ctx := c.Request.Context()

	jsonRequestBody, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.Status(400)
		log.Error(errors.New("ошибка чтения тела запроса"))
		return
	}

	if err := json.Unmarshal(jsonRequestBody, &s); err != nil {
		c.Status(400)
		c.Writer.Write([]byte("убедитесь, что вы ввели корректные данные"))
		log.Error(errors.New("ошибка декодирования json: "), err)
		return
	}

	err = sc.service.StatusController(ctx, s.TransactionID, s.Confirm)
	if err != nil {
		log.Error(err)
		return
	}

	c.Status(200)
	c.Writer.Write([]byte("вы завершили транзакцию!"))
}
