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
	posSets, err := strconv.ParseBool(c.Query("possible-sets"))
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
	action := c.Query("action")
	posSets, err := strconv.ParseBool(c.Query("possible-sets"))
	if err != nil {
		posSets = possibleSets
	}

	err = setCore.LoadState(state)
	if err != nil {
		return helpers.ThrowError(c, fiber.StatusBadRequest, err.Error())
	}

	switch action {
	case "add":
		err := setCore.AddCards()
		if err != nil {
			return helpers.ThrowError(c, fiber.StatusBadRequest, err.Error())
		}
	case "remove":
		cards := helpers.GetQueryIntSlice(c, "cards", set.SetSize)
		if len(cards) == set.SetSize {
			err := setCore.RemoveSet(cards)
			if err != nil {
				return helpers.ThrowError(c, fiber.StatusBadRequest, err.Error())
			}
		}
	}

	if posSets {
		setCore.FindSets()
	}

	return c.JSON(setCore)
}
