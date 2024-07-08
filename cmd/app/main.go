package main

import (
	v1 "finance_manager/internal/api/v1"

	"github.com/gin-gonic/gin"
)

func main() {
	// log.Print("cooking config...")

	// cfg := config.GetConfig()

	// fmt.Println(cfg)

	r := gin.Default()
	v1.Router()
	r.Run(":8080") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

	// log.Print("cooking logger...")
}
