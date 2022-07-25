package routes

import (
	"github.com/cheatsnake/shadify/internal/handlers"
	"github.com/gofiber/fiber/v2"
)

func WordsearchRouter(app fiber.Router) {
	prefix := "/api/wordsearch"

	app.Get(prefix+"/generator", handlers.WordsearchGenerator)
}
