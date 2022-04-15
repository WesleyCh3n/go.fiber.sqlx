package routes

import (
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
	"github.com/golang-jwt/jwt/v4"
	"go.fiber.sqlx/app/controllers"
	"go.fiber.sqlx/pkg/middleware"
)

func PublicRoute(app *fiber.App) {
	app.Post("/login", middleware.Login)
	app.Get("/test", restrict)
	app.Use(jwtware.New(jwtware.Config{
		SigningKey: []byte("secret"),
	}))

	api := app.Group("/api")

	api.Get("/hw", controllers.GetHws)
	api.Get("/hw/:id", controllers.GetHw)

	api.Get("/file", controllers.GetFiles)
	api.Get("/file/:id", controllers.GetFile)
}

func restrict(c *fiber.Ctx) error {
	bearToken := c.Get("Authorization")
	onlyToken := strings.Split(bearToken, " ")
	token, err := jwt.Parse(onlyToken[1], func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})
	if err != nil {
		return err
	}
	claims := token.Claims.(jwt.MapClaims)
	fmt.Println(claims["name"])

	// user := c.Locals("user").(*jwt.Token)
	// claims := user.Claims.(jwt.MapClaims)
	// name := claims["name"].(string)
	return c.SendString("Welcome " + "")
}
