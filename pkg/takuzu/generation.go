package takuzu

import (
	"errors"
	"math"
	"strconv"
	"strings"

	"github.com/cheatsnake/shadify/pkg/assists"
)

func generateField(size int) ([][]string, error) {
	var (
		index int
		row   string
	)
	field := []string{}
	currentColumns := make([]string, size)
	freeRows := generateRows(size)

	for i := 0; i < size; i++ {
		if (i > 1) {
			nextRowPattern := defineNextRow(currentColumns)
			filteredRows := filteringRows(freeRows, nextRowPattern)
			if len(filteredRows) < 1 {
				return [][]string{}, errors.New(deadEnd)
			}
			index = assists.GetRandomInteger(0, len(filteredRows) - 1)
			row = filteredRows[index]
			index = assists.IndexOf(row, freeRows)
		} else {
			index = assists.GetRandomInteger(0, len(freeRows) - 1)
			row = freeRows[index]
		}

		field = append(field, row)

		for j := 0; j < len(row); j++ {
			currentColumns[j] += string(row[j])
		}

		if index == 0 {
			freeRows = freeRows[1:]
		} else if index == len(freeRows) - 1 {
			freeRows = freeRows[:len(freeRows)-1]
		} else {
			freeRows = append(freeRows[:index], freeRows[index + 1:]...)
		}
	}
	return splitField(field), nil
}

func generateRows(size int) []string {
	rows := []string{}
	max := math.Pow(2, float64(size))

	for i := 0; i < int(max); i++ {
		binary := strconv.FormatInt(int64(i), 2)
		strBinary := assists.PadStart(binary, size - len(binary), "0")

		if strings.Contains(strBinary, strings.Repeat(takuzuZero, 3)) ||
		   strings.Contains(strBinary, strings.Repeat(takuzuOne, 3)) ||
		   strings.Count(strBinary, takuzuZero) > size / 2 ||
		   strings.Count(strBinary, takuzuOne) > size / 2 {
			continue
		}
		rows = append(rows, strBinary)
	}
	return rows
}

func defineNextRow(columns []string) string {
	size := len(columns)
	nextRow := ""

	for _, column := range(columns) {
		if column[len(column)-2:] == strings.Repeat(takuzuZero, 2) {
			nextRow += takuzuOne
		} else if column[len(column)-2:] == strings.Repeat(takuzuOne, 2) {
			nextRow += takuzuZero
		} else if strings.Count(column, takuzuZero) == size / 2 {
			nextRow += takuzuOne
		} else if strings.Count(column, takuzuOne) == size / 2 {
			nextRow += takuzuZero
		} else if len(column) < size - 2  && strings.Count(column, takuzuZero) == (size / 2) - 1 {
			nextRow += takuzuOne
		} else if len(column) < size - 2 && strings.Count(column, takuzuOne) == (size / 2) - 1 {
			nextRow += takuzuZero
		} else {
			nextRow += takuzuNull
		}
	}
	return nextRow
}

func filteringRows(rows []string, pattern string) []string {
	filteredRows := []string{}
	
	for _, row := range(rows) {
		isNextRow := true
		for i := 0; i < len(row); i++ {
			if string(pattern[i]) == takuzuNull {
				continue
			}
			if row[i] != pattern[i] {
				isNextRow = false
				break
			}
		}
		if isNextRow {
			filteredRows = append(filteredRows, row)
		}
	}
	return filteredRows
}

func splitField(field []string) [][]string {
	splittedField := [][]string{}
	for i, row := range(field) {
		splittedRow := []string{}
		for j := range(row) {
			splittedRow = append(splittedRow, string(field[i][j]))
		}
		splittedField = append(splittedField, splittedRow)
	}
	return splittedField
}

func prepareField(field [][]string, fillFactor int) ([][]string, error) {
	if len(field) == 0 {
		return [][]string{}, errors.New(emptyField)
	}
	
	preparedField := [][]string{}
	for i, row := range(field) {
		preparedRow := []string{}
		for j := range(row) {
			randBool := assists.GetRandomBool(fillFactor)
			if randBool {
				preparedRow = append(preparedRow, string(field[i][j]))
			} else {
				preparedRow = append(preparedRow, takuzuNull)
			}
		}
		preparedField = append(preparedField, preparedRow)
	}
	return preparedField, nil
}