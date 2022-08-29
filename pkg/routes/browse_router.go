package routes

import "github.com/gofiber/fiber/v2"

func browseRouter(r fiber.Router) {
	router := r.Group("/browser")

	router.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Filter top video latest, filter by type, category, country, year, time!")
	})
}
