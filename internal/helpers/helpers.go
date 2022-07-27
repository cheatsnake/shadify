package helpers

import (
	"strconv"
	"strings"

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

func GetQueryIntSlice(c *fiber.Ctx, param string, size int) []int {
	value := c.Query(param)
	res := make([]int, size)
	for i, v := range strings.Split(value, "-") {
		num, err := strconv.Atoi(v)
		if err != nil || i == size {
			return []int{}
		}
		res[i] = num
	}
	return res
}

func ThrowError(c *fiber.Ctx, code int, error string) error {
	return c.Status(code).JSON(fiber.NewError(code, error))
}
