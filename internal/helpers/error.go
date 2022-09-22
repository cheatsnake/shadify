package helpers

import "github.com/gofiber/fiber/v2"

func ThrowError(c *fiber.Ctx, code int, error string) error {
	return c.Status(code).JSON(fiber.NewError(code, error))
}
