package helpers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func GetQueryInt(c *fiber.Ctx, param string, def int) int {
	value := c.Query(param)
	number, err := strconv.Atoi(value)
	if err != nil {
		return def
	}

	return number
}

func GetQueryString(c *fiber.Ctx, param string, def string) string {
	value := c.Query(param)
	if value == "" {
		return def
	}

	return value
}