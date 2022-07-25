package minesweeper

// Core struct of Minesweeper module
type Core struct {
	StartPosition string     `json:"start"`  // Position in Core.Board that will never contain mines in and around it
	Width         int        `json:"width"`  // Width of the Core.Board
	Height        int        `json:"height"` // Height of the Core.Board
	Board         [][]string `json:"board"`  // Board with cells that contain mines and numbers corresponding to the number of mines around
	Mines         int        `json:"mines"`  // Total number of mines in the Core.Board
}
