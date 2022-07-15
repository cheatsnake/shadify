package wordsearch

import (
	"math/rand"
)

var words [25]string = [25]string{"parish", "paucity", "prim", "plasm", "hogan", "under", "ate", "butt", "why", "fay", "tar", "max", "mph", "gangster", "guaranty", "huckster", "fancy", "nolo", "chatham", "nun", "hawk", "ryan", "much", "elm", "pup"}
var alphabet [26]string = [26]string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
var directions [8][2]int = [8][2]int{{0, 1}, {1, 1}, {1, 0}, {1, -1}, {0, -1}, {-1, -1}, {-1, 0}, {-1, 1}}
var gg int = 0

func Generate(w, h, wc int) [][]string {
	grid := make([][]string, h)

	for i := range grid {
		row := make([]string, w)
		grid[i] = row
	}

	pasteWords(grid, words[:])

	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == "" { 
				grid[i][j] = alphabet[rand.Intn(26)]
			}
		}
	}

	println(gg)

	return grid
}

func pasteWords(grid [][]string, words []string) {
	var x, y, d int

	for _, word := range words {
	
		restart:
		x = rand.Intn(len(grid[0]))
		y = rand.Intn(len(grid))
		d = rand.Intn(len(directions))
	
		if isPossibleToPaste(grid, x, y, &d, word) {
			// println(x, y)
			for _, letter := range word {
				grid[y][x] = string(letter)
				x += directions[d][0]
				y += directions[d][1]
			}
			continue
		}
		goto restart
	}
}

func isPossibleToPaste(grid [][]string, x, y int, d *int, word string) bool {
	gg++
	copyX, copyY := x, y
	isPossible, checkAllDirs := true, false
	initDir, lastDir := *d,  *d - 1
	if initDir - 1 < 0 {
		lastDir = len(directions) - 1
	}
	again:
	if *d == lastDir {
		checkAllDirs = true
	}
	for _, letter := range word {
		if (copyX < len(grid[0]) && 
			copyX >= 0 &&
			copyY < len(grid) && 
			copyY >= 0 &&
			((grid[copyY][copyX] == string(letter) || grid[copyY][copyX] == ""))) {
			copyX += directions[*d][0]
			copyY += directions[*d][1]
		} else {
			if *d == len(directions) - 1 {
				*d = 0
			} else {
				*d++
			}

			if checkAllDirs { 
				isPossible = false
				break;
			} 
			copyX, copyY = x, y
			goto again
		}
	}
	// println(x, y)
	return isPossible
}