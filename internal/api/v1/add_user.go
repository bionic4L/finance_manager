package v1

import (
	"encoding/json"
	"finance_manager/internal/models"
	"io"
	"log"

	"github.com/gin-gonic/gin"
)

func addUser(c *gin.Context) {
	var u models.User

	jsonRequestBody, err := io.ReadAll(c.Request.Body)
	if err != nil {
		log.Print(err)
		return
	}

	if err := json.Unmarshal(jsonRequestBody, &u); err != nil {
		log.Print(err)
		return
	}

}
