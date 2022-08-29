package routes

import "github.com/gofiber/fiber/v2"

func watchRouter(r fiber.Router) {
	router := r.Group("/watch")

	router.Get("/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		return c.JSON(fiber.Map{
			"message": "Watching movie:" + id,
		})
	})

}
