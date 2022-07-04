package minesweeper

import (
	"strconv"

	"github.com/cheatsnake/shadify/pkg/assists"
)

func generator(sp string, w, h, mines int) [][]string {
	board := make([][]string, h)

	for i := 0; i < h; i++ {
		board[i] = make([]string, w)
	}

	for i := 0; i < mines; i++ {
		x := assists.GetRandomInteger(0, w-1)
		y := assists.GetRandomInteger(0, h-1)
		if board[y][x] == "x" {
			i--
		} else {
			board[y][x] = "x"
		}
	}

	for i := range board {
		for j := range board[i] {
			if board[i][j] != "x" {
				board[i][j] = lookAround(board, i, j)
			}
		}
	}

	return board
}

func lookAround(board [][]string, y, x int) string {
	around := 0
	for i := y - 1; i <= y+1; i++ {
		for j := x - 1; j <= x+1; j++ {
			if i < 0 || i >= len(board) || j < 0 || j >= len(board[0]) {
				continue
			}
			if board[i][j] == "x" {
				around++
			}
		}
	}
	if around == 0 {
		return "o"
	} else {
		return strconv.Itoa(around)
	}
}