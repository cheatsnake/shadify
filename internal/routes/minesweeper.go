package routes

import (
	"github.com/cheatsnake/shadify/internal/handlers"
	"github.com/gofiber/fiber/v2"
)

func MinesweeperRouter(app fiber.Router) {
	prefix := "/api/minesweeper"

	app.Get(prefix + "/generator", handlers.MinesweeperGenerator)
}