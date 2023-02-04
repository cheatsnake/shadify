package memory

import (
	"math/rand"
	"strings"
)

func gridGenerator(w, h, pairSize int) ([][]string, error) {
	letters, err := lettersPool(w*h, pairSize)
	if err != nil {
		return nil, err
	}

	grid := make([][]string, h)

	for i := range grid {
		grid[i] = letters[i*w : w*(i+1)]
	}

	return grid, nil
}

func lettersPool(totalLetters, pairSize int) ([]string, error) {
	if totalLetters < 0 {
		return nil, errNegativeNumbers
	}

	_, ok := _pairSizes[pairSize]
	if !ok {
		return nil, errPairSizeNotAllowed
	}

	if totalLetters > len(_lowerCaseLetters)*2*pairSize {
		return nil, errTooManyLetters
	}

	if totalLetters%pairSize != 0 {
		return nil, errIncorrectLettersAmount
	}

	totalPairs := totalLetters / pairSize
	allLetters := _lowerCaseLetters

	if totalPairs > len(_lowerCaseLetters) {
		allLetters += _upperCaseLetters
	}

	letters := strings.Split(strings.Repeat(allLetters[:totalPairs], pairSize), "")

	rand.Shuffle(len(letters), func(i, j int) {
		letters[i], letters[j] = letters[j], letters[i]
	})

	return letters, nil
}
