package controllers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	_ "github.com/jackc/pgx/v4/stdlib"
	"go.fiber.sqlx/platform/database"
)

func GetHws(c *fiber.Ctx) error {
	db, err := database.OpenDBConn()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	hws, err := db.GetHws()
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
			"count": 0,
			"hw":    nil,
		})
	}

	return c.JSON(fiber.Map{
		"error": false,
		"msg":   nil,
		"count": len(hws),
		"hw":    hws,
	})
}

func GetHw(c *fiber.Ctx) error {
	fmt.Fprintf(c, "%s\n", c.Params("id"))
	db, err := database.OpenDBConn()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	hw, err := db.GetHw(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
			"hw":    nil,
		})
	}

	return c.JSON(fiber.Map{
		"error": false,
		"msg":   nil,
		"hw":    hw,
	})
}
