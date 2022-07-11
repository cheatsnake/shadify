package routes

import (
	"github.com/cheatsnake/shadify/internal/handlers"
	"github.com/gofiber/fiber/v2"
)

func SchulteRouter(app fiber.Router) {
	prefix := "/api/schulte"

	app.Get(prefix + "/generator", handlers.SchulteGenerator)
}