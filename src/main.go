package main

import (
	"log"

	"github.com/Ashkan4472/google_form_go/src/pkg/config"
	"github.com/gofiber/fiber/v2"
)

func initialize() {
	config.InitialEnv()
	config.InitialDatabase()
}

func main() {
	initialize()

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	log.Fatal(app.Listen(":3000"))
}
