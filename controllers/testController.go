package controllers

import "github.com/gofiber/fiber/v2"

func Test(ctx *fiber.Ctx) error {
	return ctx.SendString("Hello, World!")
}
