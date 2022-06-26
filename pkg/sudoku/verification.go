package sudoku

import (
	"math"
	"sort"
	"strconv"
)

func verify(grid [9][9]int) *VerificationResult {
	verifyRows := checkRows(grid)
	if verifyRows.IsError {
		return verifyRows
	}

	verifyColumns := checkColumns(grid)
	if verifyColumns.IsError {
		return verifyColumns
	}

	return checkAreas(grid)
} 

func сheckLines(grid [9][9]int, lineType string) *VerificationResult {
	for i, row := range(grid) {
		sort.Ints(row[:])

		for j := 1; j <= len(row); j++ {
			if j != row[j - 1] {
				return &VerificationResult{
					true, lineType + "-" + strconv.Itoa(i + 1),
				}
			}
		}
	}
	return &VerificationResult{false, ""}
}

func checkRows(grid [9][9]int) *VerificationResult {
	return сheckLines(grid, "row")
}

func checkColumns(grid [9][9]int) *VerificationResult {
	grid = swapRowsToColumns(grid)
	return сheckLines(grid, "col")
}

func checkAreas(grid [9][9]int) *VerificationResult {
	var array [9][9]int
	areas := make([][]int, 9)

	for i := 0; i < 9; i += 3 {
		for j := 0; j < 9; j++ {
			for k := 0 + i; k < 3 + i; k++ {
				index := int(math.Floor(float64(j / 3 + i)))
				areas[index] = append(areas[index], grid[j][k])
			}
		}
	}

	for i, row := range(areas) {
		copy(array[i][:], row[:9])
	}

	return сheckLines(array, "area")
}
