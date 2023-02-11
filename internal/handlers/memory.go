package handlers

import (
	"github.com/cheatsnake/shadify/internal/helpers"
	"github.com/cheatsnake/shadify/pkg/memory"
	"github.com/gofiber/fiber/v2"
)

const (
	memoryHeight        int  = 4
	memoryWidth         int  = 6
	memoryPairSize      int  = 3
	memoryShowPositions bool = true
)

func MemoryGenerator(c *fiber.Ctx) error {
	width := helpers.GetQueryInt(c, "width", memoryWidth)
	height := helpers.GetQueryInt(c, "height", memoryHeight)
	pairSize := helpers.GetQueryInt(c, "pair-size", memoryPairSize)
	showPositions := helpers.GetQueryBool(c, "show-positions", memoryShowPositions)

	result, err := memory.Generate(width, height, pairSize, showPositions)
	if err != nil {
		return helpers.ThrowError(c, fiber.StatusBadRequest, err.Error())
	}

	return c.JSON(result)
}
