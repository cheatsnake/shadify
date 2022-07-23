package wordsearch

// Start and end position on hidden word in the *Core.Grid*
type WordPosition struct {
	Start [2]int `json:"start"` // Satrt position of the hidden word in the Core.Grid, which has the form of an array [x, y]
	End   [2]int `json:"end"`   // End position of the hidden word in the Core.Grid, which has the form of an array [x, y]
}

// Random word with it's position in Core.Grid
type Word struct {
	Word     string       `json:"word"`     // Value of hidden word
	Position WordPosition `json:"position"` // Position of the hidden word in the Core.Grid
}

// Core struct of wordsearch module
type Core struct {
	Width      int        `json:"width"`      // Width of Core.Grid
	Height     int        `json:"height"`     // Height of Core.Grid
	WordsCount int        `json:"wordsCount"` // Total words count that hidden in Core.Grid
	Grid       [][]string `json:"grid"`       // Grid consisting of letters among which are hidden random words
	Words      []Word     `json:"words"`      // Words that are hidden in Core.Grid
}
