package routes

import (
	"github.com/cheatsnake/shadify/internal/handlers"
	"github.com/gofiber/fiber/v2"
)

func TakuzuRouter(app fiber.Router) {
	prefix := "/takuzu"

	app.Get(prefix + "/generator", handlers.TakuzuGenerator)
	app.Get(prefix + "/verifier", handlers.TakuzuVerificationGet)
	app.Post(prefix + "/verifier", handlers.TakuzuVerificationPost)
}