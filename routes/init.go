package routers

import (
	"github.com/gofiber/fiber/v2"
)

func RouterInit(ctx *fiber.App) {
	Api(ctx)
}
