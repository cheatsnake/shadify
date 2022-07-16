package wordsearch

type Core struct{}

type WordPosition struct {
	Start [2]int `json:"start"`
	End   [2]int `json:"end"`
}

type Word struct {
	Word     string       `json:"word"`
	Position WordPosition `json:"position"`
}

type Wsearch struct {
	Width      int        `json:"width"`
	Height     int        `json:"height"`
	WordsCount int        `json:"wordsCount"`
	Grid       [][]string `json:"grid"`
	Words      []Word     `json:"words"`
}