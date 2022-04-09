package main

import (
	"Wedding.com/database"

	"Wedding.com/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	database.Connect()

	app := fiber.New()

	routes.Setup(app)
	app.Listen("127.0.0.1:3000")

}
