package routes

import "github.com/gofiber/fiber/v2"

func searchRouter(r fiber.Router) {
	router := r.Group("/search")

	router.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Search video by name vietnamese or english!")
	})
}
