package handlers

import (
	_ "fmt"
	"regexp"
	"strings"

	"github.com/cheatsnake/shadify/internal/helpers"
	"github.com/cheatsnake/shadify/pkg/takuzu"
	"github.com/gofiber/fiber/v2"
)

const (
	takuzuSize int = 8
	takuzuFill int = 33
)

type takuzuTaskBody struct {
	Task [][]string `json:"task"`
}

var takuzuCore *takuzu.Core

func TakuzuGenerator(c *fiber.Ctx) error {
	size := helpers.GetQueryInt(c, "size", takuzuSize)
	fill := helpers.GetQueryInt(c, "fill", takuzuFill)

	takuzuCore = takuzu.NewCore(size)
	takuzuCore.Generate()
	takuzuCore.Prepare(fill)

	return c.JSON(takuzuCore)
}

func TakuzuVerificationPost(c *fiber.Ctx) error {
	tBody := new(takuzuTaskBody)

	if err := c.BodyParser(tBody); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	takuzuCore = takuzu.NewCore(takuzuSize)
	takuzuCore.Task = tBody.Task
	result, _:= takuzuCore.Verify()

	return c.JSON(result)
}

func TakuzuVerificationGet(c *fiber.Ctx) error {
	taskStr := c.Query("task")
	matched, _ := regexp.MatchString(`[01-]`, taskStr)

	if taskStr == "" || !strings.Contains(taskStr, "-") || !matched {
		return fiber.NewError(fiber.StatusBadRequest, "invalid 'task' parameter value")
	}

	task := [][]string{}
	splitTaskStr := strings.Split(taskStr, "-")
	size := len(splitTaskStr[0])
	if size % 2 != 0 || size > 12 {
		return fiber.NewError(fiber.StatusBadRequest, "row size should be even and should not exceed 12 characters")
	}
			
	for _, row := range(splitTaskStr) {
		if len(row) == size {
			splitRow := strings.Split(row, "")
			task = append(task, splitRow)
		} else {
			return fiber.NewError(fiber.StatusBadRequest, "each row must be the same length")
		}
	}

	takuzuCore = takuzu.NewCore(size)
	takuzuCore.Task = task
	result, _ := takuzuCore.Verify()
	
	return c.JSON(result)
}
