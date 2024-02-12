package routers

import (
	"github.com/DaffaAhmadSM/CBTapp/controllers"
	"github.com/gofiber/fiber/v2"
)

func Api(app fiber.Router) {

	api := app.Group("/api")
	api.Get("/", controllers.Test)
	api.Get("/api", func(c *fiber.Ctx) error {
		return c.SendString("Hello, Api!")
	})

}
