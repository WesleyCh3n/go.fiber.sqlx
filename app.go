package main

import (
	"github.com/gofiber/fiber/v2"
	_ "github.com/joho/godotenv/autoload"
	"go.fiber.sqlx/pkg/routes"
)

func main() {
	app := fiber.New()
	routes.PublicRoute(app)
	app.Listen(":3000")
}
