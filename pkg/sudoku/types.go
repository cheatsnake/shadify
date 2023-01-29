package sudoku

type Core struct {
	Grid [9][9]int `json:"grid"`
	Task [9][9]int `json:"task"`
}

type VerificationResult struct {
	IsValid  bool   `json:"isValid"`
	Position string `json:"position"`
}
