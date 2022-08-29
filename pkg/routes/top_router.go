package routes

import "github.com/gofiber/fiber/v2"

func topRouter(r fiber.Router) {
	router := r.Group("/top")

	router.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Filter top video most views, filter by day, week month, year!")
	})
}
