package kuromasu

import (
	"math"
	"strconv"

	"github.com/cheatsnake/shadify/internal/helpers"
)

func generator(w, h int) Core {
	grid := make([][]string, h)
	fullGrid := make([][]string, h)
	boxCellsPool := int(math.Floor(float64(w) * float64(h) * boxCellPercentage))

	for i := range grid {
		row := make([]string, w)
		for j := range row {
			row[j] = emptyCell
		}
		grid[i] = row
	}

	for boxCellsPool > 0 {
		for i, row := range grid {
			for j := range row {
				if ((i-1 >= 0 && grid[i-1][j] != boxCell) || i-1 < 0) &&
					((i+1 < len(grid) && grid[i+1][j] != boxCell) || i+1 >= len(grid)) &&
					((j-1 >= 0 && row[j-1] != boxCell) || j-1 < 0) &&
					((j+1 < len(grid[0]) && row[j+1] != boxCell) || j+1 >= len(grid[0])) &&
					helpers.GetRandomBool(500) {
					row[j] = boxCell
					boxCellsPool--
				}
			}
		}
	}

	for i := range grid {
		r := make([]string, w)
		for j := 0; j < w; j++ {
			cell := boxCell
			if grid[i][j] != boxCell {
				cell = getCellViewCount(j, i, &grid)
			}
			r[j] = cell
		}
		fullGrid[i] = r
	}

	result := Core{Width: w, Height: h, Solution: fullGrid}
	return result
}

func getCellViewCount(x, y int, grid *[][]string) string {
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
