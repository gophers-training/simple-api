package handler

import "github.com/gofiber/fiber/v2"

func PrintHelloWorld(context *fiber.Ctx) error {
	return context.SendString("Hello, World!")
}

func PrintName(context *fiber.Ctx) error {
	return context.SendString("Mahmoud Ahmadi")
}
