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
	Name     string
	ImageURL string
	Desc     string
}

func main() {
	planets := []Planet{
		{
			"Mercury",
			"https://thesolarsystem.fandom.com/wiki/Mercury",
			"Closest rock to the sun",
		},
		{
			"Venus",
			"https://upload.wikimedia.org/wikipedia/commons/5/54/Venus_-_December_23_2016.png",
			"Second closest rock to the sun",
		},
		{
			"Mars",
			"https://upload.wikimedia.org/wikipedia/commons/0/02/OSIRIS_Mars_true_color.jpg",
			"Third closest rock to the sun",
		},
		{
			"Jupiter",
			"https://upload.wikimedia.org/wikipedia/commons/2/2b/Jupiter_and_its_shrunken_Great_Red_Spot.jpg",
			"Fourth closest rock to the sun",
		},
		{
			"Saturn",
			"https://upload.wikimedia.org/wikipedia/commons/c/c7/Saturn_during_Equinox.jpg",
			"Fifth closest rock to the sun",
		},
		{
			"Uranus",
			"https://upload.wikimedia.org/wikipedia/commons/3/3d/Uranus2.jpg",
			"Sixth closest rock to the sun",
		},
		{
			"Neptune",
			"https://upload.wikimedia.org/wikipedia/commons/5/56/Neptune_Full.jpg",
			"Seventh closest rock to the sun",
		},
	}
	// END: ed8c6549bwf9

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
			"planets": planets}, "layouts/main")

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
			"planet": planets,
		})
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
					"How much time do you have?",
					[]Choice{{"Less than a week", 0}, {"A week", 1}, {"More than a week", 2}},
				},
				{
					"Do you have any heart problems?",
					[]Choice{{"Yes", 0}, {"No", 1}},
				},
				{
					"Which would you prefer?",
					[]Choice{{"Sight seeing from space", 0}, {"Seeing a planet from it's surface", 1}},
				},
			},
		})
	})

	app.Post("/questionnaire", func(c *fiber.Ctx) error {
		log.Print(c.Request())
		return c.Render("pages/homepage", fiber.Map{})
	})

	app.Get("/elements", func(c *fiber.Ctx) error {
		// Render index within layouts/main
		return c.Render("pages/homepage", fiber.Map{
			"planet": planets,
		})
	})

	log.Fatal(app.Listen(":3000"))
}
