package main

import (
	"log"

	routesv1 "github.com/Ashkan4472/google_form_go/src/internals/routes/routes_v1"
	"github.com/Ashkan4472/google_form_go/src/pkg/config"
	"github.com/gofiber/fiber/v2"
)

func initialize() {
	config.InitialEnv()
	config.InitialDatabase()
}

func setupRoutes(app *fiber.App) {
	r := app.Group("api/v1")
	routesv1.SetupAuthRoutes(r)
}

func main() {
	initialize()

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})
	setupRoutes(app)
	log.Fatal(app.Listen(":3000"))
}
