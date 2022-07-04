package minesweeper

func NewCore() *Core {
	return &Core{}
}

func (sc *Core) Generate(sp string, w, h, mines int) Board {
	board := generator(sp, w, h, mines)

	return Board{
		StartPosition: sp,
		Width:         w,
		Height:        h,
		Board:         board,
		Mines:         mines,
	}
}
