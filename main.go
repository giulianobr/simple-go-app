package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

func setupApp() *fiber.App {
	// Initialize standard Go html template engine
	engine := html.New("./web/html", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Static("/resources", "./web/resources")

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", nil)
	})

	return app
}

func main() {
	app := setupApp()
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	log.Fatal(app.Listen(":" + port))
}
