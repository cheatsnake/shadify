package minesweeper

import "errors"

func NewCore() *Core {
	return &Core{}
}

func (sc *Core) Generate(sp string, w, h, mines int) (Board, error) {

	if mines >= (w * h) {
		return Board{}, errors.New("the field is too small for so many mines")
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
