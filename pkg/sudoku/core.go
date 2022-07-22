package sudoku

import (
	"github.com/cheatsnake/shadify/pkg/assists"
)

// Create a new instance of Sudoku Core
func NewCore() *Core {
	return &Core{}
}

// Generate a new Sudoku 9x9 grid 
func (sc *Core) Generate() {
	sc.Grid = basicGrid
	numberOfSwaps := assists.GetRandomInteger(10, 50)

	for i := 0; i < numberOfSwaps; i++ {
		sc.Grid = swapRows(sc.Grid)
		sc.Grid = swapColumns(sc.Grid)
		sc.Grid = swapRowsToColumns(sc.Grid)
		sc.Grid = swapRowAreas(sc.Grid)
		sc.Grid = swapColAreas(sc.Grid)
	}
}

// Prepare a new Sudoku task based on already generated grid
func (sc *Core) Prepare(fillFactor int) {
	sc.Task = prepare(sc.Grid, fillFactor)
}

// Verify a current Sudoku task
func (sc *Core) Verify() *VerificationResult {
	return verify(sc.Task)
}