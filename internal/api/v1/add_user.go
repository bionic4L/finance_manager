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

type CreateUser struct {
	service *service.CreateUserService
}

func CreateUserRouter(r *gin.Engine, service *service.CreateUserService) {
	cu := &CreateUser{service: service}
	r.POST("/user-add", cu.addUser)
}

func (cu *CreateUser) addUser(c *gin.Context) {
	ctx := c.Request.Context()

	var u *models.User
	jsonRequestBody, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.Status(400)
		log.Error(errors.New("ошибка чтения тела запроса"))
		return
	}

	if err := json.Unmarshal(jsonRequestBody, &u); err != nil {
		c.Status(400)
		c.Writer.Write([]byte("убедитесь, что вы ввели корректные данные"))
		log.Error(errors.New("ошибка декодирования json"))
		return
	}

	if err := ValidateAddUser(c, u); err != nil {
		c.Status(400)
		c.Writer.Write([]byte("валидация запроса не пройдена"))
		log.Warn(err)
		return
	}

	if err := cu.service.UserCreate(ctx, u.Name); err != nil {
		log.Error(err)
		c.Status(400)
		return
	}

	c.Status(200)
	c.JSON(200, u.Name)
	c.Writer.Write([]byte("пользователь создан!"))

}

func ValidateAddUser(c *gin.Context, u *models.User) error {
	if u.Name == "" {
		return errors.New("пустое поле имени пользователя")
	}

	if len(u.Name) < 4 {
		return errors.New("слишком короткое имя пользователя (минимум 4 символа)")
	}
	return nil
}
