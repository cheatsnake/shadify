package handlers

import (
	"github.com/cheatsnake/shadify/internal/helpers"
	"github.com/cheatsnake/shadify/pkg/countries"
	"github.com/gofiber/fiber/v2"
)

const countriesVariants = 4

func CountriesCountryQuiz(c *fiber.Ctx) error {
	variants := helpers.GetQueryInt(c, "variants", countriesVariants)
	cq, err := countries.GetCountryQuiz(variants)
	if err != nil {
		return helpers.ThrowError(c, fiber.StatusBadRequest, err.Error())
	}

	return c.JSON(cq)
}

func CountriesCapitalQuiz(c *fiber.Ctx) error {
	variants := helpers.GetQueryInt(c, "variants", countriesVariants)
	cq, err := countries.GetCapitalQuiz(variants)
	if err != nil {
		return helpers.ThrowError(c, fiber.StatusBadRequest, err.Error())
	}

	return c.JSON(cq)
}
