package main

import (
	"fmt"
	"golang-rest-api/internal/app"
	"golang-rest-api/internal/config"
)

func main() {
	cfg := config.LoadConfig(".env")

	application := app.NewApp(cfg)

	if err := application.Run(); err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}
