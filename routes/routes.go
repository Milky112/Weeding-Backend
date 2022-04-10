package routes

import (
	"fmt"

	"Wedding.com/controllers"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	fmt.Println("Routes is calling on docker images")

	app.Get("/api/data/:user_data", controllers.GetAllData)

	app.Post("/api/comment/", controllers.PostComment)

}
