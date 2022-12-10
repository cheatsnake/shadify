package kuromasu

import (
	"math"
	"strconv"

	"github.com/cheatsnake/shadify/internal/helpers"
)

func generator(w, h, fill int) Core {
	grid := make([][]string, h)
	boxPool := generateBoxPool(w, h)

	for i := range grid {
		row := make([]string, w)

		for j := range row {
			row[j] = emptyCell
		}

		grid[i] = row
	}

	for boxPool > 0 {
		x := helpers.GetRandomInteger(0, w-1)
		y := helpers.GetRandomInteger(0, h-1)

		checkUp := (y+1 < len(grid) && grid[y+1][x] != boxCell) || y+1 >= len(grid)
		checkDown := (y-1 >= 0 && grid[y-1][x] != boxCell) || y-1 < 0
		checkRight := (x+1 < len(grid[0]) && grid[y][x+1] != boxCell) || x+1 >= len(grid[0])
		checkLeft := (x-1 >= 0 && grid[y][x-1] != boxCell) || x-1 < 0
		couldBeBox := helpers.GetRandomBool(500)

		if checkUp && checkDown && checkRight && checkLeft && couldBeBox {
			grid[y][x] = boxCell
			boxPool--
		}
	}

	for i := range grid {
		for j := 0; j < w; j++ {
			cell := boxCell

			if grid[i][j] != boxCell {
				cell = calcViewCount(j, i, &grid)
			}

			grid[i][j] = cell
		}
	}

	task := prepareTask(grid, fill)

	result := Core{Width: w, Height: h, Solution: grid, Task: task}
	return result
}

func generateBoxPool(w, h int) int {
	randBoxPercent := float64(helpers.GetRandomInteger(minBoxPercent, maxBoxPercent)) / 100
	boxPool := int(math.Round(float64(w) * float64(h) * randBoxPercent))
	return boxPool
}

func calcViewCount(x, y int, grid *[][]string) string {
	count := 1

	for i := 0; i < cellDirections; i++ {
		xx := x
		yy := y

		for xx >= 0 && yy >= 0 &&
			xx < len((*grid)[0]) &&
			yy < len(*grid) &&
			(*grid)[yy][xx] != boxCell {
			if !(xx == x && yy == y) {
				count++
			}

			if i == 0 {
				yy++
			} else if i == 1 {
				xx++
			} else if i == 2 {
				yy--
			} else {
				xx--
			}
		}

		if i == 3 {
			break
		}
	}

	return strconv.Itoa(count)
}

func prepareTask(sol [][]string, fill int) [][]string {
	task := make([][]string, len(sol))
	totalCells := len(sol) * len(sol[0])
	fillCount := int(math.Round(float64(totalCells) * float64(fill) / 100))

	for i := range task {
		row := make([]string, len(sol[0]))

		for j := range row {
			row[j] = emptyCell
		}

		task[i] = row
	}

	for fillCount > 0 {
		x := helpers.GetRandomInteger(0, len(sol[0])-1)
		y := helpers.GetRandomInteger(0, len(sol)-1)

		if sol[y][x] != boxCell {
			task[y][x] = sol[y][x]
			fillCount--
		}
	}

	return task
}
