package v1

import (
	"encoding/json"
	"finance_manager/internal/models"
	"finance_manager/internal/service"
	"io"
	"log"

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
	var u models.User

	jsonRequestBody, err := io.ReadAll(c.Request.Body)
	if err != nil {
		log.Print(err)
		return
	}

	if err := json.Unmarshal(jsonRequestBody, &u.Name); err != nil {
		log.Print(err)
		return
	}

	if err := cu.service.CreateUser(u.Name); err != nil {
		log.Print(err)
		c.Status(400)
		return
	}

	c.Status(200)
	c.JSON(200, u.Name)

}
