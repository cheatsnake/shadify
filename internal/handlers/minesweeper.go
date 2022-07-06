package handlers

import (
	"github.com/cheatsnake/shadify/internal/helpers"
	"github.com/cheatsnake/shadify/pkg/minesweeper"
	"github.com/gofiber/fiber/v2"
)

const (
	minesweeperWidth  = 9
	minesweeperHeight = 9
	minesweeperMines  = 12
)

var minesweeperCore *minesweeper.Core

func init() {
	minesweeperCore = minesweeper.NewCore()
}

func MinesweeperGenerator(c *fiber.Ctx) error {
	start := helpers.GetQueryString(c, "start", "")
	width := helpers.GetQueryInt(c, "width", minesweeperWidth)
	height := helpers.GetQueryInt(c, "height", minesweeperHeight)
	mines := helpers.GetQueryInt(c, "mines", minesweeperMines)

	result, err := minesweeperCore.Generate(start, width, height, mines)
	if err != nil {
		code := fiber.StatusBadRequest
		return c.Status(code).JSON(fiber.NewError(code, err.Error()))
	}

	return c.JSON(result)
}