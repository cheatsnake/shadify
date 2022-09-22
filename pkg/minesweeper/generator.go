package minesweeper

import (
	"errors"
	"strconv"
	"strings"

	"github.com/cheatsnake/shadify/internal/helpers"
)

func generator(startX, startY, w, h, mines int) [][]string {
	board := make([][]string, h)

	for i := 0; i < h; i++ {
		board[i] = make([]string, w)
	}

	for i := 0; i < mines; i++ {
		x := helpers.GetRandomInteger(0, w-1)
		y := helpers.GetRandomInteger(0, h-1)
		if board[y][x] == "x" || isProtectedPosition(x, y, startX, startY) {
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

func parseStartPosition(sp string, w, h int) (int, int, error) {
	posValues := strings.Split(sp, "-")
	if len(posValues) < 2 {
		return 0, 0, errors.New(incorrectStartPos)
	}

	x, errX := strconv.Atoi(posValues[0])
	y, errY := strconv.Atoi(posValues[1])

	if errX != nil || errY != nil || x > w || y > h || x < 1 || y < 1 {
		return 0, 0, errors.New(incorrectStartPos)
	}

	return x - 1, y - 1, nil
}

func isProtectedPosition(x, y, startX, startY int) bool {
	if x == startX && y == startY {
		return true
	}

	for i := y - 1; i <= y+1; i++ {
		for j := x - 1; j <= x+1; j++ {
			if i == startY && j == startX {
				return true
			}
		}
	}

	return false
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
