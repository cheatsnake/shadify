package routes

import (
	"github.com/cheatsnake/shadify/internal/handlers"
	"github.com/gofiber/fiber/v2"
)

func CampRouter(app fiber.Router) {
	prefix := "/api/camp"

	app.Get(prefix+"/generator", handlers.CampGenerator)
	app.Post(prefix+"/verifier", handlers.CampVerifier)
}
