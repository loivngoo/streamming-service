package routes

import "github.com/gofiber/fiber/v2"

func SetupRouterV1(r *fiber.App) {

	v1 := r.Group("/api/v1")

	browseRouter(v1)
	searchRouter(v1)
	topRouter(v1)
	movieRouter(v1)
	tvRouter(v1)
	watchRouter(v1)
}
