package minesweeper

import "errors"

func NewCore() *Core {
	return &Core{}
}

func (sc *Core) Generate(sp string, w, h, mines int) (Board, error) {

	if (w * h < minFieldCells) {
		return Board{}, errors.New(tooSmallField)
	}

	if (w * h > maxFieldCells) {
		return Board{}, errors.New(tooBigField)
	}

	if float64(mines) > float64(w * h) * (maxMinesPercent) {
		return Board{}, errors.New(tooManyMines)
	}

	startX, startY, err := parseStartPosition(sp, w, h)
	if err != nil {
		return Board{}, err
	}

	board := generator(startX, startY, w, h, mines)

	return Board{
		StartPosition: sp,
		Width:         w,
		Height:        h,
		Board:         board,
		Mines:         mines,
	}, nil
}
