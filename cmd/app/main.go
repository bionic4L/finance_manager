package main

import (
	"finance_manager/internal/app"

	log "github.com/sirupsen/logrus"
)

func main() {
	if err := app.Run("config/local.yaml"); err != nil {
		log.Fatal(err)
	}
}
