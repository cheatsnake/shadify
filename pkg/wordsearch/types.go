package wordsearch

type WordPosition struct {
	Start [2]int
	End   [2]int
}

type Word struct {
	Word     string
	Position WordPosition
}

type Wsearch struct {
	Width      int
	Height     int
	WordsCount int
	Grid       [][]string
	Words      []Word
}