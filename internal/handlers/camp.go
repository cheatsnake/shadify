package handlers

import (
	"github.com/cheatsnake/shadify/internal/helpers"
	"github.com/cheatsnake/shadify/pkg/camp"
	"github.com/gofiber/fiber/v2"
)

const (
	campWidth    int  = 7
	campHeight   int  = 7
	campSolution bool = true
)

type campTaskBody struct {
	Solution    [][]int `json:"solution"`
	RowTents    []int   `json:"rowTents"`
	ColumnTents []int   `json:"columnTents"`
}

func CampGenerator(c *fiber.Ctx) error {
	width := helpers.GetQueryInt(c, "width", campWidth)
	height := helpers.GetQueryInt(c, "height", campHeight)
	solution := helpers.GetQueryBool(c, "solution", campSolution)

	result, err := camp.Generate(width, height)
	if err != nil {
		return helpers.ThrowError(c, fiber.StatusBadRequest, err.Error())
	}

	if !solution {
		result.Solution = [][]int{}
	}

	return c.JSON(result)
}

func CampVerifier(c *fiber.Ctx) error {
	tBody := new(campTaskBody)

	if err := c.BodyParser(tBody); err != nil {
		return helpers.ThrowError(c, fiber.StatusBadRequest, err.Error())
	}

	result, err := camp.Verify(tBody.Solution, tBody.RowTents, tBody.ColumnTents)
	if err != nil {
		return helpers.ThrowError(c, fiber.StatusBadRequest, err.Error())
	}

	return c.JSON(result)
}
