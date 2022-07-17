package main

import (
	"log"
	"os"
	"time"

	"github.com/cheatsnake/shadify/internal/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
)

func main() {
	app := fiber.New()
	port := os.Getenv("PORT")

	if os.Getenv("PORT") == "" {
		port = "5000"
	}

	app.Use(limiter.New(limiter.Config{
		Expiration: 10 * time.Second,
		Max: 5,
	}))

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Redirect("https://github.com/cheatsnake/shadify/blob/master/README.md#documentation")
	})

	routes.MathRouter(app)
	routes.TakuzuRouter(app)
	routes.SudokuRouter(app)
	routes.SchulteRouter(app)
	routes.SetRouter(app)
	routes.MinesweeperRouter(app)
	routes.WordsearchRouter(app)

	log.Fatal(app.Listen(":" + port))
}