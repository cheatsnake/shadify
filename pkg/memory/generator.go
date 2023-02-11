package memory

import (
	"math/rand"
	"strings"
)

func generator(w, h, pairSize int, showPositions bool) ([][]string, []PairPositions, error) {
	if w < 0 || h < 0 {
		return nil, nil, errNegativeNumbers
	}

	pool, usedLetters, err := lettersPool(w*h, pairSize)
	if err != nil {
		return nil, nil, err
	}

	grid := make([][]string, h)

	for i := range grid {
		grid[i] = pool[i*w : w*(i+1)]
	}

	if showPositions {
		pp := definePositions(grid, usedLetters, pairSize)
		return grid, pp, nil
	}

	return grid, nil, nil
}

func lettersPool(totalLetters, pairSize int) ([]string, string, error) {
	_, ok := _pairSizes[pairSize]
	if !ok {
		return nil, "", errPairSizeNotAllowed
	}

	if totalLetters > len(_lowerCaseLetters)*2*pairSize {
		return nil, "", errTooBigGrid
	}

	if totalLetters%pairSize != 0 {
		return nil, "", errIncorrectGridSize
	}

	totalPairs := totalLetters / pairSize
	allLetters := _lowerCaseLetters

	if totalPairs > len(_lowerCaseLetters) {
		allLetters += _upperCaseLetters
	}

	usedLetters := allLetters[:totalPairs]
	pool := strings.Split(strings.Repeat(usedLetters, pairSize), "")

	rand.Shuffle(len(pool), func(i, j int) {
		pool[i], pool[j] = pool[j], pool[i]
	})

	return pool, usedLetters, nil
}

func definePositions(grid [][]string, letters string, ps int) []PairPositions {
	values := make(map[string][][]int)

	for i, row := range grid {
		for j, val := range row {
			_, ok := values[val]
			if !ok {
				values[val] = make([][]int, 0, ps)
			}

			values[val] = append(values[val], []int{i + 1, j + 1})
		}
	}

	pp := make([]PairPositions, len(letters))

	for i, letter := range letters {
		pp[i] = PairPositions{
			Value:     string(letter),
			Positions: values[string(letter)],
		}
	}

	return pp
}
