package anagram

import (
	"path/filepath"

	"github.com/cheatsnake/shadify/internal/helpers"
)

var wordsDB []string
var wordsDBSize int

func init() {
	nouns, _ := helpers.ReadFileLineByLine(filepath.Join("data", "nouns.txt"))
	wordsDB = nouns
	wordsDBSize = len(nouns)
}

// Generate an anagram with a random task and find possible words
func Generate() Core {
	word := getRandomWord(9, 15)
	words := composingWords(word)

	return Core{Task: word, Words: words}
}
