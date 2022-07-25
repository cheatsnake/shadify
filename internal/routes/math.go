package routes

import (
	"github.com/cheatsnake/shadify/internal/handlers"
	"github.com/gofiber/fiber/v2"
)

func MathRouter(app fiber.Router) {
	prefix := "/api/math"

	app.Get(prefix+"/add", handlers.MathAddition)
	app.Get(prefix+"/sub", handlers.MathSubtraction)
	app.Get(prefix+"/mul", handlers.MathMultiplication)
	app.Get(prefix+"/div", handlers.MathDivision)
	app.Get(prefix+"/quad", handlers.MathQuadratic)
}
