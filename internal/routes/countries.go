package routes

import (
	"github.com/cheatsnake/shadify/internal/handlers"
	"github.com/gofiber/fiber/v2"
)

func CountriesRouter(app fiber.Router) {
	prefix := "/api/countries"

	app.Get(prefix+"/country-quiz", handlers.CountriesCountryQuiz)
	app.Get(prefix+"/capital-quiz", handlers.CountriesCapitalQuiz)
}
