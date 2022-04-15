package main

import (
	"fmt"

	"Wedding.com/database"

	Logger "Wedding.com/log"
	"Wedding.com/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	database.Connect()
	fmt.Println("Running Docker Apps")

	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
		AllowOrigins:     "*",
		AllowHeaders:     "Origin, Content-Type, Accept, Accept-Language, Content-Length",
	}))

	Logger.GenerateLog()
	Logger.CommonLog.Println("Server is Running")
	app.Use(logger.New())

	routes.Setup(app)
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!") // send text
	})
	app.Listen(":3000")

}
