package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()
	app.Use(cors.New())

	// app.Get("/fetch/*/:where?", Fetch)
	// app.Get("/create/:name", Create)
	// app.Post("/update/:name", Update)
	// app.Post("/config/update/:name", UpdateConfig)

	app.Get("/fetch/:what?/:where?", Fetch)
	app.Get("/gatekeeper/:who", Gatekeeper)
	app.Post("/update/:what/:where", Update)
	app.Post("/register/:what/:where", Register)

	app.Listen(":8080")
}
