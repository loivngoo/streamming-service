package routes

import "github.com/gofiber/fiber/v2"

func tvRouter(r fiber.Router) {

	router := r.Group("/tv")
	{
		router.Get("/", func(c *fiber.Ctx) error {
			return c.SendString("API get list TVShows!")
		})

		router.Post("/", func(c *fiber.Ctx) error {
			return c.SendString("API add a TVShow!")
		})

		router.Delete("/", func(c *fiber.Ctx) error {
			return c.SendString("API remove a TVShow!")
		})

		router.Patch("/", func(c *fiber.Ctx) error {
			return c.SendString("API update a TVShow!")
		})

		router.Get("/:id", func(c *fiber.Ctx) error {
			id := c.Params("id")
			return c.JSON(fiber.Map{
				"id":   id,
				"name": "TVShow " + id,
			})
		})

	}

}
