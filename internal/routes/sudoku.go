package routes

import (
	"github.com/cheatsnake/shadify/internal/handlers"
	"github.com/gofiber/fiber/v2"
)

func SudokuRouter(app fiber.Router) {
	prefix := "/sudoku"

	app.Get(prefix + "/generator", handlers.SudokuGenerator)
	app.Get(prefix + "/verify", handlers.SudokuVerificationGet)
	app.Post(prefix + "/verify", handlers.SudokuVerificationPost)
}