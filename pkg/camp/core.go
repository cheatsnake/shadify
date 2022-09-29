package camp

import "errors"

func Generate(w, h int) (Core, error) {
	if w > maxFieldSize || h > maxFieldSize {
		return Core{}, errors.New(tooLargeFieldSize)
	}

	if w < minFieldSize || h < minFieldSize {
		return Core{}, errors.New(tooSmallFieldSize)
	}

	trees := calcTreesCount(w, h)
	field := generateField(w, h, trees)
	rowTents, columnTents := countTents(field)
	task := hideTents(field)

	return Core{w, h, trees, rowTents, columnTents, task, field}, nil
}

func Verify(field [][]int, rowTents, columnTents []int) (VerifyResult, error) {
	if len(field) > maxFieldSize || len(field[0]) > maxFieldSize {
		return VerifyResult{}, errors.New(tooLargeFieldSize)
	}

	if len(field) < minFieldSize || len(field[0]) < minFieldSize {
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
