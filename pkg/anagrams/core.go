package anagrams

import (
	"path/filepath"

	"github.com/cheatsnake/shadify/pkg/assists"
)

var wordsDB []string
var wordsDBSize int

func init() {
	nouns, err := assists.ReadFileLineByLine(filepath.Join("data", "nouns.txt"))
	if err != nil {
		fmt.Println(err)
	}
	wordsDB = nouns
	wordsDBSize = len(nouns)
}

// Generate an anagram with random task and random hiden words
func Generate() Anagram {
	word := getRandomWord(9, 15)
	words := composingWords(word)

	return Anagram{Task: word, Words: words}
}