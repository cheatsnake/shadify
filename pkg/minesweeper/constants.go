package minesweeper

const (
	maxMinesPercent float64 = 0.25
	minFieldCells   int     = 16
	maxFieldCells   int     = 1000
)

// error constants
const (
	tooSmallField     string = "field is too small"
	tooBigField       string = "field is too big"
	tooManyMines      string = "the field is too small for so many mines"
	incorrectStartPos string = "incorrect start position"
)
