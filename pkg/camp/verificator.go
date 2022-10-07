package camp

import "fmt"

func verifyTask(field [][]int, rowTents, columnTents []int) VerifyResult {
	for i, row := range field {
		if calcTents(row) != rowTents[i] {
			return VerifyResult{true, fmt.Sprintf("row-%d", i+1), wrongRowTents}
		}
	}

	for i := 0; i < len(field[0]); i++ {
		column := make([]int, 0, len(columnTents))
		for j := 0; j < len(field); j++ {
			column = append(column, field[j][i])
		}
		if calcTents(column) != columnTents[i] {
			return VerifyResult{true, fmt.Sprintf("col-%d", i+1), wrongColumnTents}
		}
	}

	for i := 0; i < len(field); i++ {
		for j := 0; j < len(field[0]); j++ {
			if field[i][j] == tentValue {
				field[i][j] = emptyValue
				result := checkIsAreaFit(j, i, field, true, false)
				field[i][j] = tentValue

				if !result {
					return VerifyResult{
						true,
						fmt.Sprintf("row-%d, col-%d", i+1, j+1),
						wrongTentPosition}
				}
			}
		}
	}

	return VerifyResult{false, "", ""}
}

func calcTents(values []int) int {
	result := 0

	for _, val := range values {
		if val == tentValue {
			result++
		}
	}

	return result
}
