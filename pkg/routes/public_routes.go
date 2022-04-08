package routes

import (
	"github.com/gofiber/fiber/v2"
	"go.fiber.sqlx/app/controllers"
)

func PublicRoute(app *fiber.App) {
	api := app.Group("/api")

	api.Get("/hw", controllers.GetHws)
	api.Get("/hw/:id", controllers.GetHw)

	api.Get("/file", controllers.GetFiles)
	api.Get("/file/:id", controllers.GetFile)
}
