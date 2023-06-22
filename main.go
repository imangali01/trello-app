package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"

	"github.com/Imangali2002/trello-app/config"
	"github.com/Imangali2002/trello-app/handler"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("[-] No .env file found")
		os.Exit(1)
	}
}

func main() {
	config := config.GetConfig()

	app := &handler.App{}
	app.Initialize(config)
	app.Run(":8080")
}
