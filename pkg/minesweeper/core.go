package minesweeper

import "errors"

// Generate a new Core with given start position, width
// and height of the Core.Board and total number of mines
func Generate(sp string, w, h, mines int) (Core, error) {

	if (w * h < minFieldCells) {
		return Core{}, errors.New(tooSmallField)
	}

	if (w * h > maxFieldCells) {
		return Core{}, errors.New(tooBigField)
	}

	if float64(mines) > float64(w * h) * (maxMinesPercent) {
		return Core{}, errors.New(tooManyMines)
	}

	startX, startY, err := parseStartPosition(sp, w, h)
	if err != nil {
		return Core{}, err
	}

	board := generator(startX, startY, w, h, mines)

	return Core{
		StartPosition: sp,
		Width:         w,
		Height:        h,
		Board:         board,
		Mines:         mines,
	}, nil
}
