package routes

import (
	"Wedding.com/controllers"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	app.Get("/api/data/:user_data", controllers.GetAllData)

	app.Post("/api/comment/", controllers.PostComment)

}
