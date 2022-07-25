package routes

import (
	"github.com/cheatsnake/shadify/internal/handlers"
	"github.com/gofiber/fiber/v2"
)

func AnagramRouter(app fiber.Router) {
	prefix := "/api/anagram"

	app.Get(prefix+"/generator", handlers.AnagramGenerator)
}
