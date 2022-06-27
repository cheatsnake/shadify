package main

import (
	"log"
	"os"
	_ "time"

	"github.com/cheatsnake/shadify/internal/routes"
	"github.com/gofiber/fiber/v2"
	_ "github.com/gofiber/fiber/v2/middleware/limiter"
)

func main() {
	app := fiber.New()
	port := os.Getenv("PORT")

	if os.Getenv("PORT") == "" {
		port = "5000"
	}

	// app.Use(limiter.New(limiter.Config{
	// 	Expiration: 10 * time.Second,
	// 	Max: 5,
	// }))

	routes.MathRouter(app)
	routes.TakuzuRouter(app)
	routes.SudokuRouter(app)
	routes.SchulteRouter(app)
	routes.SetRouter(app)

	log.Fatal(app.Listen(":" + port))
}