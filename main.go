package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/django/v3"
)

type Question struct {
	Text    string
	Choices []Choice
}

type Choice struct {
	Text string
	Id   int
}

type Planet struct {
	Name string
	Desc string
}

func main() {
	// Create a new engine
	engine := django.New("./views", ".django")

	// Or from an embedded system
	// See github.com/gofiber/embed for examples
	// engine := html.NewFileSystem(http.Dir("./views", ".django"))

	// Pass the engine to the Views
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Get("/", func(c *fiber.Ctx) error {
		// Render index
		return c.Render("index", fiber.Map{
			"planets": []Planet{
				{
					"Mercury",
					"Closest rock to the sun",
				},
				{
					"Venus",
					"Second closest rock to the sun",
				},
			},
		}, "layouts/main")

	})

	app.Get("/layout", func(c *fiber.Ctx) error {
		// Render index within layouts/main
		return c.Render("index", fiber.Map{
			"Title": "Hello, World!",
		}, "layouts/main")
	})

	app.Get("/homepage", func(c *fiber.Ctx) error {
		// Render index within layouts/main
		return c.Render("pages/homepage", fiber.Map{
			"planets": []Planet{
				{
					"Mercury",
					"Closest rock to the sun",
				},
				{
					"Venus",
					"Second closest rock to the sun",
				},
			},
		})
	})

	app.Get("/questionnaire", func(c *fiber.Ctx) error {
		// Render index within layouts/main
		return c.Render("pages/questionnaire", fiber.Map{})
	})

	app.Get("/destination", func(c *fiber.Ctx) error {
		// Render index within layouts/main
		return c.Render("pages/destination", fiber.Map{})
	})

	log.Fatal(app.Listen(":3000"))
}
