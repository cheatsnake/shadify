package sudoku

import "github.com/cheatsnake/shadify/internal/helpers"

func prepare(grid [9][9]int, fillFactor int) [9][9]int {
	if fillFactor > 99 {
		return grid
	}
	if fillFactor < 0 {
		fillFactor = 0
	}

	filled := 0
	task := getEmptyGrid()

	for filled < fillFactor {
		randRow := helpers.GetRandomInteger(0, 8)
		randCol := helpers.GetRandomInteger(0, 8)

		if task[randRow][randCol] == 0 {
			task[randRow][randCol] = grid[randRow][randCol]
			filled++
		}
	}

	return task
}

func getEmptyGrid() [9][9]int {
	grid := [9][9]int{}
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid); j++ {
			grid[i][j] = 0
		}
	}
	return grid
}
