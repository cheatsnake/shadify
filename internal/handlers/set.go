package handlers

import (
	"strconv"

	"github.com/cheatsnake/shadify/internal/helpers"
	"github.com/cheatsnake/shadify/pkg/set"
	"github.com/gofiber/fiber/v2"
)

const (
	possibleSets = true
)

var setCore *set.Core

func init() {
	setCore = set.NewCore()
}

func SetGameStart(c *fiber.Ctx) error {
	posSets, err := strconv.ParseBool(c.Query("possibleSets"))
	if err != nil {
		posSets = possibleSets
	}

	setCore.Start()

	if posSets {
		setCore.FindSets()
	}

	return c.JSON(setCore)
}

func SetGameAllCards(c *fiber.Ctx) error {
	return c.JSON(set.Deck)
}

func SetGameLoadState(c *fiber.Ctx) error {
	state := c.Params("state")
	rSet := helpers.GetQueryIntSlice(c, "removeSet", set.SetSize)
	add, err := strconv.ParseBool(c.Query("addCards"))
	if err != nil {
		add = false
	}

	posSets, err := strconv.ParseBool(c.Query("possibleSets"))
	if err != nil {
		posSets = possibleSets
	}

	err = setCore.LoadState(state)
	if err != nil {
		code := fiber.StatusBadRequest
		return c.Status(code).JSON(fiber.NewError(code, err.Error()))
	}

	if posSets {
		setCore.FindSets()
	}

	if len(rSet) == set.SetSize {
		err := setCore.RemoveSet(rSet)
		if err != nil {
			code := fiber.StatusBadRequest
			return c.Status(code).JSON(fiber.NewError(code, err.Error()))
		} else {
			if posSets {
				setCore.FindSets()
			}
			return c.JSON(setCore)
		}
	}

	if add {
		err := setCore.AddCards()
		if err != nil {
			code := fiber.StatusBadRequest
			return c.Status(code).JSON(fiber.NewError(code, err.Error()))
		} else {
			if posSets {
				setCore.FindSets()
			}
			return c.JSON(setCore)
		}
	}

	return c.JSON(setCore)
}

