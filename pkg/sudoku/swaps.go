package sudoku

import (
	"math"

	"github.com/cheatsnake/shadify/internal/helpers"
)

func swapRowsToColumns(grid [9][9]int) [9][9]int {
	swappedGrid := [9][9]int{}
	for i, row := range grid {
		for j, elem := range row {
			swappedGrid[j][i] = elem
		}
	}
	return swappedGrid
}

func swapRows(grid [9][9]int) [9][9]int {
	areaSize := int(math.Sqrt(float64(len(grid))))
	area := helpers.GetRandomInteger(0, areaSize-1)
	first := helpers.GetRandomInteger(0, areaSize-1)
	firstRow := area*areaSize + first

	second := helpers.GetRandomInteger(0, areaSize-1)
	isEqual := first == second

	for isEqual {
		second = helpers.GetRandomInteger(0, areaSize-1)
		isEqual = first == second
	}

	secondRow := area*areaSize + second

	grid[firstRow], grid[secondRow] = grid[secondRow], grid[firstRow]

	return grid
}

func swapColumns(grid [9][9]int) [9][9]int {
	grid = swapRowsToColumns(grid)
	grid = swapRows(grid)
	grid = swapRowsToColumns(grid)
	return grid
}

func swapRowAreas(grid [9][9]int) [9][9]int {
	areaSize := int(math.Sqrt(float64(len(grid))))
	fisrtArea := helpers.GetRandomInteger(0, areaSize-1)
	secondArea := helpers.GetRandomInteger(0, areaSize-1)
	isEqual := fisrtArea == secondArea

	for isEqual {
		secondArea = helpers.GetRandomInteger(0, areaSize-1)
		isEqual = fisrtArea == secondArea
	}

	for i := 0; i < areaSize; i++ {
		firstIndex, secondIndex := fisrtArea*areaSize+i, secondArea*areaSize+i
		grid[firstIndex], grid[secondIndex] = grid[secondIndex], grid[firstIndex]
	}

	return grid
}

func swapColAreas(grid [9][9]int) [9][9]int {
	grid = swapRowsToColumns(grid)
	grid = swapRowAreas(grid)
	grid = swapRowsToColumns(grid)
	return grid
}
