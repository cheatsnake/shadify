package routes

import (
	"github.com/cheatsnake/shadify/internal/handlers"
	"github.com/gofiber/fiber/v2"
)

func SetRouter(app fiber.Router) {
	prefix := "/set"

	app.Get(prefix + "/start", handlers.SetGameStart)
	app.Get(prefix + "/cards", handlers.SetGameAllCards)
	app.Get(prefix + "/state/:state", handlers.SetGameLoadState)
}
