package main

import (
	"api-gateway/internal/app"
	"api-gateway/internal/pkg/config"
)

func main() {
	cf := config.Load()
	app.Run(cf)
}
