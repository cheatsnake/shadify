package handlers

import (
	"math"

	"github.com/cheatsnake/shadify/internal/helpers"
	"github.com/cheatsnake/shadify/pkg/schulte"
	"github.com/gofiber/fiber/v2"
)

const (
	schulteSize int    = 5
	schulteMode string = "numbers"
)

func SchulteGenerator(c *fiber.Ctx) error {
	size := helpers.GetQueryInt(c, "size", schulteSize)
	mode := helpers.GetQueryString(c, "mode", schulteMode)

	if mode == "alphabet" {
		result := schulte.GenerateAlphabetic()
		return c.JSON(result)
	}

	if math.Abs(float64(size)) > 15 {
		return helpers.ThrowError(c, fiber.StatusBadRequest, "size should not exceed 15")
	}

	result := schulte.GenerateNumeric(size)
	return c.JSON(result)
}
