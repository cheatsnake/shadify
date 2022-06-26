package handlers

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/cheatsnake/shadify/internal/helpers"
	"github.com/cheatsnake/shadify/pkg/sudoku"
	"github.com/gofiber/fiber/v2"
)

const (
	sudokuFill int = 30
)

type sudokuTaskBody struct {
	Task [9][9]int `json:"task"`
}

var sudokuCore *sudoku.Core

func SudokuGenerator(c *fiber.Ctx) error {
	fill := helpers.GetQueryInt(c, "fill", sudokuFill)

	sudokuCore = sudoku.NewCore()
	sudokuCore.Generate()
	sudokuCore.Prepare(fill)

	return c.JSON(sudokuCore)
}

func SudokuVerificationPost(c *fiber.Ctx) error {
	tBody := new(sudokuTaskBody)

	if err := c.BodyParser(tBody); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	sudokuCore = sudoku.NewCore()
	sudokuCore.Task = tBody.Task
	result := sudokuCore.Verify()

	return c.JSON(result)
}

func SudokuVerificationGet(c *fiber.Ctx) error {
	taskStr := c.Query("task")
	matched, _ := regexp.MatchString(`[1-9-]`, taskStr)

	if taskStr == "" || !strings.Contains(taskStr, "-") || !matched {
		return fiber.NewError(fiber.StatusBadRequest, "invalid 'task' parameter value")
	}

	sudokuCore = sudoku.NewCore()
	task := make([][]int, 9)
	splitedTaskStr := strings.Split(taskStr, "-")
			
	for i, row := range(splitedTaskStr) {
		if len(row) == 9 {
			splitRow := strings.Split(row, "")
			sudokuRow := []int{}
			for j := 0; j < 9; j++ {
				number, _ := strconv.Atoi(splitRow[j])
				sudokuRow = append(sudokuRow, number)
			}

			task[i] = sudokuRow
			copy(sudokuCore.Task[i][:], task[i][0:9])
		} else {
			return fiber.NewError(fiber.StatusBadRequest, "each row must be 9 digits long")
		}
	}

	result := sudokuCore.Verify()
	
	return c.JSON(result)
}