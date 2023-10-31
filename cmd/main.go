package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/benebobaa/batik-api-go/database"
)

func main() {
	database.ConnectDb()
	
	app := fiber.New()

	setupRoutes(app)

	log.Fatal(app.Listen(":3000"))
}
