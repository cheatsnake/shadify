package handlers

import (
	"math"

	"github.com/cheatsnake/shadify/internal/helpers"
	"github.com/cheatsnake/shadify/pkg/schulte"
	"github.com/gofiber/fiber/v2"
)

type schulteGrid[T int | string] struct {
	Grid [][]T `json:"grid"`
}

const (
	schulteSize int = 5
	schulteMode string = "numbers"
)

var schulteCore *schulte.Core

func init() {
	schulteCore = schulte.NewCore()
}

func SchulteGenerator(c *fiber.Ctx) error {
	size := helpers.GetQueryInt(c, "size", schulteSize)
	mode := helpers.GetQueryString(c, "mode", schulteMode)

	if mode == "alphabet" {
		grid := schulteGrid[string]{Grid: schulteCore.GenerateAlphabetic()}
		return c.JSON(grid)
	}

	if (math.Abs(float64(size)) > 15) {
		code := fiber.StatusBadRequest
		return c.Status(code).JSON(fiber.NewError(code, "size should not exceed 15"))
	}

	grid := schulteGrid[int]{Grid: schulteCore.GenerateNumeric(size)}
	return c.JSON(grid)
}