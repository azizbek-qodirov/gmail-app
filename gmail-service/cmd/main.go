package main

import (
	"gmail-service/internal/app"
	"gmail-service/internal/pkg/config"
)

func main() {
	config := config.Load()

	app.Run(&config)
}
