package routes

import (
	"github.com/cheatsnake/shadify/internal/handlers"
	"github.com/gofiber/fiber/v2"
)

func KuromasuRouter(app fiber.Router) {
	prefix := "/api/kuromasu"

	app.Get(prefix+"/generator", handlers.KuromasuGenerator)
	app.Post(prefix+"/verifier", handlers.KuromasuVerifier)
}
