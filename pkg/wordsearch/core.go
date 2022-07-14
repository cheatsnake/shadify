package wordsearch

import (
	"math/rand"
	"time"
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


	// for _, w := range words {
	// 	pasteWord(grid, w)
	// }
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
	rand.Seed(time.Now().UnixNano())
	var x, y, d int

	for _, word := range words {
	
		restart:
		x = rand.Intn(len(grid[0]))
		y = rand.Intn(len(grid))
		d = rand.Intn(len(directions))
	
		if isPossibleToPaste(grid, x, y, d, word) {
			println(x, y)
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

func isPossibleToPaste(grid [][]string, x, y, d int, word string) bool {
	gg++
	isPossible := true
	// println(x, y)
	for _, letter := range word {
		if (x < len(grid[0]) && 
			x >= 0 &&
			y < len(grid) && 
			y >= 0 &&
			((grid[y][x] == string(letter) || grid[y][x] == ""))) {
			x += directions[d][0]
			y += directions[d][1]
		} else {
			isPossible = false
			break;
		}
	}
	// println(x, y)
	return isPossible
}