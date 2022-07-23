package handlers

import (
	"github.com/cheatsnake/shadify/internal/helpers"
	"github.com/cheatsnake/shadify/pkg/wordsearch"
	"github.com/gofiber/fiber/v2"
)

const (
	wordsearchGridWidth  = 9
	wordsearchGridHeight = 9
)

func WordsearchGenerator(c *fiber.Ctx) error {
	width := helpers.GetQueryInt(c, "width", wordsearchGridWidth)
	height := helpers.GetQueryInt(c, "height", wordsearchGridHeight)

	wsc, err := wordsearch.Generate(width, height)
	if err != nil {
		code := fiber.StatusBadRequest
		return c.Status(code).JSON(fiber.NewError(code, err.Error()))
	}

	return c.JSON(wsc)
}