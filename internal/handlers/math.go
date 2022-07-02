package handlers

import (
	"github.com/cheatsnake/shadify/internal/helpers"
	"github.com/cheatsnake/shadify/pkg/math"
	"github.com/gofiber/fiber/v2"
)

const (
	mathMin int = 1
	mathMax int = 99
	mathNegative int = 0
	mathQuadMin int = 1
	mathQuadMax int = 20 
)

var mathCore *math.Core	

func init() {
	mathCore = math.NewCore()
}

func MathAddition(c *fiber.Ctx) error {
	minFirst := helpers.GetQueryInt(c, "minFirst", mathMin)
	maxFirst := helpers.GetQueryInt(c, "maxFirst", mathMax)
	minSecond := helpers.GetQueryInt(c, "minSecond", mathMin)
	maxSecond := helpers.GetQueryInt(c, "maxSecond", mathMax)

	result := mathCore.Addition(minFirst, maxFirst, minSecond, maxSecond)
	
	return c.JSON(result)
}

func MathSubtraction(c *fiber.Ctx) error {
	minFirst := helpers.GetQueryInt(c, "minFirst", mathMin)
	maxFirst := helpers.GetQueryInt(c, "maxFirst", mathMax)
	minSecond := helpers.GetQueryInt(c, "minSecond", mathMin)
	maxSecond := helpers.GetQueryInt(c, "maxSecond", mathMax)
	negative := helpers.GetQueryInt(c, "negative", mathNegative)
	allowNegative := negative == 1

	result := mathCore.Subtraction(minFirst, maxFirst, minSecond, maxSecond, allowNegative)
	
	return c.JSON(result)
}

func MathMultiplication(c *fiber.Ctx) error {
	minFirst := helpers.GetQueryInt(c, "minFirst", mathMin)
	maxFirst := helpers.GetQueryInt(c, "maxFirst", mathMax)
	minSecond := helpers.GetQueryInt(c, "minSecond", mathMin)
	maxSecond := helpers.GetQueryInt(c, "maxSecond", mathMax)

	result := mathCore.Multiplication(minFirst, maxFirst, minSecond, maxSecond)

	return c.JSON(result)
}

func MathDivision(c *fiber.Ctx) error {
	minFirst := helpers.GetQueryInt(c, "minFirst", mathMin)
	maxFirst := helpers.GetQueryInt(c, "maxFirst", mathMax)

	result := mathCore.Division(minFirst, maxFirst)

	return c.JSON(result)
}

func MathQuadratic(c *fiber.Ctx) error {
	minA := helpers.GetQueryInt(c, "minA", mathQuadMin)
	minB := helpers.GetQueryInt(c, "minB", mathQuadMin)
	minC := helpers.GetQueryInt(c, "minC", mathQuadMin)
	maxA := helpers.GetQueryInt(c, "maxA", mathQuadMax)
	maxB := helpers.GetQueryInt(c, "maxB", mathQuadMax * 2)
	maxC := helpers.GetQueryInt(c, "maxC", mathQuadMax)

	result := mathCore.Quadratic(minA, maxA, minB, maxB, minC, maxC)

	return c.JSON(result)
} 