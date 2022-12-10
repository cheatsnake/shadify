package kuromasu

import "fmt"

// app constant
const (
	emptyCell      string = "o"
	boxCell        string = "x"
	cellDirections int    = 4

	minSideLength int = 4
	maxSideLength int = 15

	minBoxPercent  int = 15
	maxBoxPercent  int = 25
	minFillPercent int = 10
	maxFillPercent int = 50
)

// error constant
var (
	failedMaxSideLength = fmt.Sprintf("length of each side must not exceed %d", maxSideLength)
	failedMinSideLength = fmt.Sprintf("length of each side must not be less than %d", minSideLength)
	failedFillPercent   = fmt.Sprintf("fill level should be between %d and %d", minFillPercent, maxFillPercent)
)
