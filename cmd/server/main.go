package main

import (
	"log"
	"os"
	"time"

	"github.com/cheatsnake/shadify/internal/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/limiter"
)

func main() {
	app := fiber.New()
	port := os.Getenv("PORT")

	if os.Getenv("PORT") == "" {
		port = "5000"
	}

	app.Static("/", "./book")

	app.Use(cors.New())
	app.Use(limiter.New(limiter.Config{
		Expiration: 10 * time.Second,
		Max:        15,
	}))

	routes.MathRouter(app)
	routes.TakuzuRouter(app)
	routes.SudokuRouter(app)
	routes.SchulteRouter(app)
	routes.SetRouter(app)
	routes.MinesweeperRouter(app)
	routes.WordsearchRouter(app)
	routes.AnagramRouter(app)
	routes.CountriesRouter(app)
	routes.CampRouter(app)
	routes.KuromasuRouter(app)
	routes.MemoryRouter(app)

	log.Fatal(app.Listen(":" + port))
}
