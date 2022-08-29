package routes

import "github.com/gofiber/fiber/v2"

func movieRouter(r fiber.Router) {
	router := r.Group("/movie")
	{
		router.Get("/", func(c *fiber.Ctx) error {
			return c.SendString("API get list movies!")
		})

		router.Post("/", func(c *fiber.Ctx) error {
			return c.SendString("API add a movie!")
		})

		router.Delete("/", func(c *fiber.Ctx) error {
			return c.SendString("API remove a movie!")
		})

		router.Patch("/", func(c *fiber.Ctx) error {
			return c.SendString("API update a movie!")
		})

		router.Get("/:id", func(c *fiber.Ctx) error {
			id := c.Params("id")
			return c.JSON(fiber.Map{
				"id":   id,
				"name": "Movie " + id,
			})
		})

	}
}
