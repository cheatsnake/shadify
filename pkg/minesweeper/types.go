package minesweeper

type Core struct{}

type Board struct {
	StartPosition string     `json:"start"`
	Width         int        `json:"width"`
	Height        int        `json:"height"`
	Board         [][]string `json:"board"`
	Mines         int        `json:"mines"`
}