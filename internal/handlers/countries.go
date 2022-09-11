package handlers

import (
	"github.com/cheatsnake/shadify/internal/helpers"
	"github.com/cheatsnake/shadify/pkg/countries"
	"github.com/gofiber/fiber/v2"
)

const (
	countriesVariants   = 4
	countriesQuizAmount = 1
)

func CountriesCountryQuiz(c *fiber.Ctx) error {
	variants := helpers.GetQueryInt(c, "variants", countriesVariants)
	amount := helpers.GetQueryInt(c, "amount", countriesQuizAmount)
	cq, err := countries.GetCountryQuiz(variants, amount)
	if err != nil {
		return helpers.ThrowError(c, fiber.StatusBadRequest, err.Error())
	}

	if amount == countriesQuizAmount {
		return c.JSON(cq[0])
	}

	return c.JSON(cq)
}

func CountriesCapitalQuiz(c *fiber.Ctx) error {
	variants := helpers.GetQueryInt(c, "variants", countriesVariants)
	amount := helpers.GetQueryInt(c, "amount", countriesQuizAmount)
	cq, err := countries.GetCapitalQuiz(variants, amount)
	if err != nil {
		return helpers.ThrowError(c, fiber.StatusBadRequest, err.Error())
	}

	if amount == countriesQuizAmount {
		return c.JSON(cq[0])
	}

	return c.JSON(cq)
}
