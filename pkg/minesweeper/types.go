package minesweeper

type Core struct{}

type Board struct {
	StartPosition string
	Width         int
	Height        int
	Board         [][]string
	Mines         int
}