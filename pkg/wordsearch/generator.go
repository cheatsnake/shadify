package wordsearch

import (
	"math"
	"math/rand"

	"github.com/cheatsnake/shadify/pkg/assists"
)

func generateGrid(w, h int) ([][]string, []Word) {
	var x, y, d int
	var pastedWords []Word
	
	restart:
	pastedWords = []Word{}
	words := getRandomWords(w, h)
	grid := make([][]string, h)
	for i := range grid {
		row := make([]string, w)
		grid[i] = row
	}

	cells := assists.GetNumbers(len(grid)*len(grid[0]), 0)

	for _, word := range words {
		changePos:
		randCell := rand.Intn(len(cells))
		y = int(math.Floor(float64(cells[randCell]) / float64(len(grid[0]))))
		x = cells[randCell] - (len(grid[0]) * y)
		cells = assists.RemoveElement(cells, randCell)
		d = rand.Intn(len(directions))

		if isPossibleToPaste(grid, x, y, &d, word) {
			startPos := [2]int{x + 1, y + 1}
			for i, letter := range word {
				grid[y][x] = string(letter)
				if (i == (len(word) - 1)) {
					break;
				}
				x += directions[d][0]
				y += directions[d][1]
			}
			
			pastedWords = append(pastedWords, Word{
				Word: word, 
				Position: 
				WordPosition{Start: startPos, End: [2]int{x + 1, y + 1}} })

			if len(cells) < 1 {
				goto restart
			}

			continue
		}

		if len(cells) < 1 {
			goto restart
		}

		goto changePos
	}
	return grid, pastedWords
}

func isPossibleToPaste(grid [][]string, x, y int, d *int, word string) bool {
	copyX, copyY := x, y
	isPossible, checkAllDirs := true, false
	initDir, lastDir := *d, *d-1

	if initDir-1 < 0 {
		lastDir = len(directions) - 1
	}

	again:
	if *d == lastDir {
		checkAllDirs = true
	}

	for _, letter := range word {
		if copyX < len(grid[0]) &&
			copyX >= 0 &&
			copyY < len(grid) &&
			copyY >= 0 &&
			(grid[copyY][copyX] == string(letter) || grid[copyY][copyX] == "") {
			copyX += directions[*d][0]
			copyY += directions[*d][1]
		} else {
			if *d == len(directions)-1 {
				*d = 0
			} else {
				*d++
			}

			if checkAllDirs {
				isPossible = false
				break
			}
			copyX, copyY = x, y
			goto again
		}
	}

	return isPossible
}