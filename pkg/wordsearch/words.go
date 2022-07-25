package wordsearch

import (
	"math"
	"math/rand"

	"github.com/cheatsnake/shadify/pkg/assists"
)

func getRandomWords(w, h int) []string {
	words := make([]string, 0, w*h/4)
	maxLettersCount := int(math.Floor(float64(w*h) * maxGridFill))
	lettersCount := 0

	for lettersCount < maxLettersCount {
		randWord := wordsDB[rand.Intn(len(wordsDB))]
		rwLen := len(randWord)

		if assists.IndexOf(randWord, words) == -1 && rwLen <= w && rwLen <= h {
			lettersCount += rwLen
			words = append(words, randWord)
		}
	}

	return words
}
