package main

import (
	routers "github.com/DaffaAhmadSM/CBTapp/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

func main() {
	fiberInit := fiber.New()
	routers.RouterInit(fiberInit)
	log.Fatal(fiberInit.Listen(":8080"))
}
