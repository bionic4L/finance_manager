package main

import (
	"finance_manager/internal/app"
	"log"
)

func main() {
	log.Fatal(app.Run("config/local.yaml"))
}
