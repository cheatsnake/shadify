package anagram

import (
	"path/filepath"

	"github.com/cheatsnake/shadify/pkg/assists"
)

var wordsDB []string
var wordsDBSize int

func init() {
	nouns, err := assists.ReadFileLineByLine(filepath.Join("data", "nouns.txt"))
	if err != nil {
		println(err.Error())
	}
	wordsDB = nouns
	wordsDBSize = len(nouns)
}

// Generate an anagram with a random task and find possible words
func Generate() Core {
	word := getRandomWord(9, 15)
	words := composingWords(word)

	return Core{Task: word, Words: words}
}