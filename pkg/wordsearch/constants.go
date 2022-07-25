package wordsearch

var alphabet [26]string = [26]string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
var directions [8][2]int = [8][2]int{{0, 1}, {1, 1}, {1, 0}, {1, -1}, {0, -1}, {-1, -1}, {-1, 0}, {-1, 1}}

const (
	minGridSideSize = 5
	maxGridSideSize = 20
	maxCellsCount   = 256
	maxGridFill     = 0.8
)

// Error constant
const (
	minGridSideError string = "the size of each side of the grid should not be less than 5 units"
	maxGridSideError string = "the size of each side of the grid must not be larger than 20 units"
	tooManyCells     string = "the total number of grid cells cannot exceed 256 units"
)
