package handlers

import (
	"github.com/cheatsnake/shadify/pkg/anagram"
	"github.com/gofiber/fiber/v2"
)

func AnagramGenerator(c *fiber.Ctx) error {
	result := anagram.Generate()
	return c.JSON(result)
}