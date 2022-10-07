package camp

import (
	"errors"
)

func Generate(w, h int) (Core, error) {
	if w > maxSideSize || h > maxSideSize {
		return Core{}, errors.New(tooLargeFieldSize)
	}

	if w < minSideSize || h < minSideSize {
		return Core{}, errors.New(tooSmallFieldSize)
	}

	trees := calcTreesCount(w, h)
	field := generateField(w, h, trees)
	rowTents, columnTents := countTents(field)
	task := hideTents(field)

	return Core{w, h, trees, rowTents, columnTents, task, field}, nil
}

func Verify(field [][]int, rowTents, columnTents []int) (VerifyResult, error) {
	if len(field) > maxSideSize || len(field[0]) > maxSideSize {
		return VerifyResult{}, errors.New(tooLargeFieldSize)
	}

	if len(field) < minSideSize || len(field[0]) < minSideSize {
		return VerifyResult{}, errors.New(tooSmallFieldSize)
	}

	if len(field) != len(rowTents) {
		return VerifyResult{}, errors.New(incorrectRowTents)
	}

	if len(field[0]) != len(columnTents) {
		return VerifyResult{}, errors.New(incorrectColumnTents)
	}

	result := verifyTask(field, rowTents, columnTents)

	return result, nil
}
