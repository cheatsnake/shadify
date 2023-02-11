package routes

import (
	"github.com/cheatsnake/shadify/internal/handlers"
	"github.com/gofiber/fiber/v2"
)

func MemoryRouter(app fiber.Router) {
	prefix := "/api/memory"

	app.Get(prefix+"/generator", handlers.MemoryGenerator)
}
