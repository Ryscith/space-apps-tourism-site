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
	Name                         string
	ImageURL                     string
	Description                  string
	Weather                      string
	Atmosphere                   string
	GravityRelativeToEarth       float32
	DayNightCycleRelativeToEarth float32
	AxisRotationAngle            int
}

func main() {
	planets := []Planet{
		{
			Name:                         "Mercury",
			ImageURL:                     "https://upload.wikimedia.org/wikipedia/commons/thumb/4/4a/Mercury_in_true_color.jpg/220px-Mercury_in_true_color.jpg",
			Description:                  "Closest rock to the sun",
			Weather:                      "No atmosphere, extremely hot on the day side and cold on the night side",
			Atmosphere:                   "Very thin exosphere with traces of hydrogen, helium, and oxygen",
			GravityRelativeToEarth:       0.38,
			DayNightCycleRelativeToEarth: 58.6,
			AxisRotationAngle:            2,
		},
		{
			Name:                         "Venus",
			ImageURL:                     "https://upload.wikimedia.org/wikipedia/commons/5/54/Venus_-_December_23_2016.png",
			Description:                  "Second closest rock to the sun",
			Weather:                      "Thick clouds, sulfuric acid rain, extremely hot surface temperatures",
			Atmosphere:                   "96.5% carbon dioxide, clouds of sulfuric acid",
			GravityRelativeToEarth:       0.91,
			DayNightCycleRelativeToEarth: 243, // Venus has a retrograde rotation
			AxisRotationAngle:            177,
		},
		{
			Name:                         "Mars",
			ImageURL:                     "https://upload.wikimedia.org/wikipedia/commons/0/02/OSIRIS_Mars_true_color.jpg",
			Description:                  "Fourth planet from the sun",
			Weather:                      "Thin atmosphere, very cold, dust storms",
			Atmosphere:                   "95.3% carbon dioxide, 2.7% nitrogen, 1.6% argon",
			GravityRelativeToEarth:       0.375,
			DayNightCycleRelativeToEarth: 1.03,
			AxisRotationAngle:            25,
		},
		{
			Name:                         "Jupiter",
			ImageURL:                     "https://upload.wikimedia.org/wikipedia/commons/2/2b/Jupiter_and_its_shrunken_Great_Red_Spot.jpg",
			Description:                  "Largest planet in the solar system",
			Weather:                      "Stormy with the Great Red Spot being a persistent storm",
			Atmosphere:                   "Predominantly hydrogen and helium",
			GravityRelativeToEarth:       2.528,
			DayNightCycleRelativeToEarth: 0.41,
			AxisRotationAngle:            3,
		},
		{
			Name:                         "Saturn",
			ImageURL:                     "https://imgs.search.brave.com/QDIf3cy7Q8445Uv6CZ0XiP1pnTJOIT1H5tRPMSnhs-U/rs:fit:860:0:0/g:ce/aHR0cHM6Ly90NC5m/dGNkbi5uZXQvanBn/LzAwLzY2Lzk0LzIx/LzM2MF9GXzY2OTQy/MTg4X1ZDSWVsTjlX/Q3g5ZlUyRGpVbmRa/aVp4aDU4bHc0VVZR/LmpwZw",
			Description:                  "Known for its stunning rings",
			Weather:                      "Wind speeds up to 1,800 km/h",
			Atmosphere:                   "Mostly hydrogen and helium",
			GravityRelativeToEarth:       1.065,
			DayNightCycleRelativeToEarth: 0.45,
			AxisRotationAngle:            27,
		},
		{
			Name:                         "Uranus",
			ImageURL:                     "https://upload.wikimedia.org/wikipedia/commons/3/3d/Uranus2.jpg",
			Description:                  "A pale blue planet tilted on its side",
			Weather:                      "Cold and cloudy",
			Atmosphere:                   "Hydrogen, helium, and methane",
			GravityRelativeToEarth:       0.886,
			DayNightCycleRelativeToEarth: 0.72,
			AxisRotationAngle:            98,
		},
		{
			Name:                         "Neptune",
			ImageURL:                     "https://upload.wikimedia.org/wikipedia/commons/5/56/Neptune_Full.jpg",
			Description:                  "Farthest planet from the sun",
			Weather:                      "Strong storms with high wind speeds",
			Atmosphere:                   "Hydrogen, helium, and methane",
			GravityRelativeToEarth:       1.14,
			DayNightCycleRelativeToEarth: 0.67,
			AxisRotationAngle:            28,
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
			"planets": planets,
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
		return c.Render("pages/results", fiber.Map{
			"planet_results": planets,
		})
	})

	app.Get("/destination/:planetname", func(c *fiber.Ctx) error {
		var target_planet Planet
		for _, planet := range planets {
			if planet.Name == c.Params("planetname") {
				target_planet = planet
			}
		}

		// Render index within layouts/main
		return c.Render("pages/destination", fiber.Map{
			"planet": target_planet,
		})
	})

	app.Get("/elements", func(c *fiber.Ctx) error {
		// Render index within layouts/main
		return c.Render("pages/homepage", fiber.Map{
			"planet": planets,
		})
	})

	app.Get("/tourpackage/:packagename", func(c *fiber.Ctx) error {
		// Render index within layouts/main
		return c.Render("pages/tourpackage", fiber.Map{})
	})

	log.Fatal(app.Listen(":3000"))
}
