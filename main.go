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
			"Title": "Hello, World!",
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
		return c.Render("pages/homepage", fiber.Map{})
	})

	app.Get("/questionnaire", func(c *fiber.Ctx) error {
		// Render index within layouts/main
		return c.Render("pages/questionnaire", fiber.Map{
			"questions": []Question{
				{
					"What temperature do you prefer?",
					[]Choice{{"Hot", 0}, {"Temperate", 1}, {"Cold", 2}, {"No Preference", 3}},
				},
				{
					"Do you like geology?",
					[]Choice{{"I love it!", 0}, {"I hate it.", 1}, {"No Preference", 2}},
				},
				{
					"Do you like geology?",
					[]Choice{{"I love it!", 0}, {"I hate it.", 1}, {"No Preference", 2}},
				},
				{
					"Do you like geology?",
					[]Choice{{"I love it!", 0}, {"I hate it.", 1}, {"No Preference", 2}},
				},
				{
					"Do you like geology?",
					[]Choice{{"I love it!", 0}, {"I hate it.", 1}, {"No Preference", 2}},
				},
			},
		})
	})

	app.Post("/questionnaire", func(c *fiber.Ctx) error {
		log.Print(c.Request())
		return c.Render("pages/homepage", fiber.Map{})
	})

	log.Fatal(app.Listen(":3000"))
}
