package sudoku

type Core struct {
	Grid [9][9]int `json:"grid"`
	Task [9][9]int `json:"task"`
}

type VerificationResult struct {
	IsError  bool   `json:"isError"`
	Position string `json:"position"`
}