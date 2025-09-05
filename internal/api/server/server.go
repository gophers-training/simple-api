package apiserver

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gophers-training/simple-api/internal/api/handler"
)

func Start() {
	app := fiber.New()

	app.Get("/print/name", handler.PrintName)
	app.Get("/print/message", handler.PrintHelloWorld)

	app.Listen(":3000")
}
