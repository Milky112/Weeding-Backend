package main

import (
	"fmt"

	"Wedding.com/database"

	"Wedding.com/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	database.Connect()
	fmt.Println("Running Docker Apps")

	app := fiber.New()

	fmt.Println("Setup Docker Routes")

	routes.Setup(app)
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!") // send text
	})
	app.Listen("127.0.0.1:3000")

}
