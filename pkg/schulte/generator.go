package schulte

import (
	"math"

	"github.com/cheatsnake/shadify/internal/helpers"
)

func generateGrid[T int | string](values []T) [][]T {
	size := int(math.Sqrt(float64(len(values))))
	available := len(values) - 1
	grid := [][]T{}

	for i := 0; i < size; i++ {
		row := make([]T, size)
		grid = append(grid, row)
	}

	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			index := helpers.GetRandomInteger(0, available)
			grid[i][j] = values[index]
			values = helpers.RemoveElement(values, index)
			available -= 1
		}
	}

	return grid
}
