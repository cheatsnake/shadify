package wordsearch

import (
	"errors"
	"math/rand"
	"path/filepath"

	"github.com/cheatsnake/shadify/pkg/assists"
)

var wordsDB []string

func init() {
	wordsDB, _ = assists.ReadFileLineByLine(filepath.Join("data", "nouns.txt"))
}

func NewCore() *Core {
	return &Core{}
}

func (wc *Core) Generate(w, h int) (Wsearch, error) {
	if w < minGridSideSize || h < minGridSideSize {
		return Wsearch{}, errors.New(minGridSideError)
	}

	if w > maxGridSideSize || h > maxGridSideSize {
		return Wsearch{}, errors.New(maxGridSideError)
	}

	if w * h > maxCellsCount {
		return Wsearch{}, errors.New(tooManyCells)
	}

	grid, pWords := generateGrid(w, h)

	// fill unused cells by random letters
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == "" { 
				grid[i][j] = alphabet[rand.Intn(26)]
			}
		}
	}

	return Wsearch{
		Width: w,
		Height: h,
		WordsCount: len(pWords),
		Grid: grid,
		Words: pWords,
	}, nil
}