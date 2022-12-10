package handlers

import (
	"github.com/cheatsnake/shadify/internal/helpers"
	"github.com/cheatsnake/shadify/pkg/kuromasu"
	"github.com/gofiber/fiber/v2"
)

const (
	kuromasuSize int = 5
	kuromasuFill int = 30
)

type kuromasuTaskBody struct {
	Task [][]string `json:"task"`
}

func KuromasuGenerator(c *fiber.Ctx) error {
	w := helpers.GetQueryInt(c, "width", kuromasuSize)
	h := helpers.GetQueryInt(c, "height", kuromasuSize)
	f := helpers.GetQueryInt(c, "fill", kuromasuFill)

	result, err := kuromasu.Generate(w, h, f)
	if err != nil {
		return helpers.ThrowError(c, fiber.StatusBadRequest, err.Error())
	}

	return c.JSON(result)
}

func KuromasuVerifier(c *fiber.Ctx) error {
	tBody := new(kuromasuTaskBody)

	if err := c.BodyParser(tBody); err != nil {
		return helpers.ThrowError(c, fiber.StatusBadRequest, err.Error())
	}

	result, err := kuromasu.Verify(tBody.Task)
	if err != nil {
		return helpers.ThrowError(c, fiber.StatusBadRequest, err.Error())
	}

	return c.JSON(result)
}