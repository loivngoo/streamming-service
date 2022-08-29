package main

import (
	"streamming-service/pkg/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {

	app := fiber.New(fiber.Config{
		AppName: "Streamming Service v1.0.1",
	})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"title": "Streamming Service",
		})
	})

	routes.SetupRouterV1(app)
	app.Listen(":8080")

}
