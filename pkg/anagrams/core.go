package anagrams

import (
	"fmt"
	"math/rand"
	"path/filepath"
	"strings"

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

func Generate() {
	word := getRandomWord(9, 15)
	words := composingWords(word)

	fmt.Println(Anagram{Task: word, Words: words})
}

func getRandomWord(minLength, maxLength int) string {
	var word string

	for len(word) < minLength || len(word) > maxLength {
		word = wordsDB[rand.Intn(wordsDBSize)]
	}

	return word
}

func composingWords(word string) []string {
	result := make([]string, 0, 32)

	for _, w := range wordsDB {
		isFits := true
		for _, letter := range w {
			if !(strings.Count(w, string(letter)) <= strings.Count(word, string(letter))) || len(w) < 3 {
				isFits = false
			}
		}

		if isFits {
			result = append(result, w)
		}
	}

	result = assists.RemoveElement(result, assists.IndexOf(word, result))

	return result
}