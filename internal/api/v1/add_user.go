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
	var u *models.User
	ctx := c.Request.Context()

	jsonRequestBody, err := io.ReadAll(c.Request.Body)
	defer c.Request.Body.Close()
	if err != nil {
		c.Status(400)
		log.Error("ошибка чтения тела запроса: ", err)
		return
	}

	if err := json.Unmarshal(jsonRequestBody, &u); err != nil {
		c.JSON(400, "проверьте правильность введенных данных")
		log.Error("ошибка декодирования json: ", err)
		return
	}

	if err := ValidateAddUser(c, u); err != nil {
		c.JSON(400, "валидация запроса не пройдена")
		log.Warn(err)
		return
	}

	if err := cu.service.UserCreate(ctx, u.Name); err != nil {
		log.Error(err)
		c.Status(400)
		return
	}
	c.JSON(200, "пользователь создан")

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
