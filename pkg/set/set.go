package set

import (
	"math/bits"
)

func findSets(cards []Card) [][]Card {
	cardsId := make([]int, len(cards))
	sets := make([][]Card, 0, DeckSize)

	for i, card := range cards {
		cardsId[i] = card.Id
	}

	length := uint(len(cardsId))

	for subsetBits := 1; subsetBits < (1 << length); subsetBits++ {
		if SetSize > 0 && bits.OnesCount(uint(subsetBits)) != SetSize {
			continue
		}

		var setId []int

		for object := uint(0); object < length; object++ {
			if (subsetBits>>object)&1 == 1 {
				setId = append(setId, cardsId[object])
			}
		}

		if validateSet(setId) {
			set := make([]Card, SetSize)
			for i, id := range setId {
				set[i] = Deck[id]
			}

			sets = append(sets, set)
		}
	}

	return sets
}

func validateSet(setId []int) bool {
	result := compareFeature(Deck[setId[0]].Number, Deck[setId[1]].Number, Deck[setId[2]].Number) &&
		compareFeature(Deck[setId[0]].Shape, Deck[setId[1]].Shape, Deck[setId[2]].Shape) &&
		compareFeature(Deck[setId[0]].Shading, Deck[setId[1]].Shading, Deck[setId[2]].Shading) &&
		compareFeature(Deck[setId[0]].Color, Deck[setId[1]].Color, Deck[setId[2]].Color)

	return result
}

func compareFeature[T int | string](f1, f2, f3 T) bool {
	return ((f1 == f2) && (f2 == f3) || (f1 != f2) && (f2 != f3) && (f3 != f1))
}
