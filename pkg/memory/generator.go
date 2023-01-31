package memory

import (
	"errors"
	"math/rand"
	"strings"
)

func generateGrid(w, h, pairSize int) {
	// grid := make([][]string, h)

	// letters, _ := createLetterPairs(w*h, pairSize)
}

func createLetterPairs(total, pairSize int) ([]string, error) {
	if total > len(_lowerCaseLetters)*2 {
		return nil, errors.New("too big size")
	}

	totalPairs := total / pairSize
	allLetters := _lowerCaseLetters

	if totalPairs > len(_lowerCaseLetters) {
		allLetters += _upperCaseLetters
	}

	letters := strings.Split(strings.Repeat(allLetters[:totalPairs], pairSize), "")
	shuffleSlice(letters)

	return letters, nil

}

func shuffleSlice(s []string) {
	rand.Shuffle(len(s), func(i, j int) {
		s[i], s[j] = s[j], s[i]
	})
}
