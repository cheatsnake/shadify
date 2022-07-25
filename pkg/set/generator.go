package set

import (
	"sort"

	"github.com/cheatsnake/shadify/pkg/assists"
)

func startGame() ([]Card, []Card) {
	cardIndexes := assists.GetNumbers(DeckSize, 0)
	layout := make([]Card, StartLayoutSize)
	freeCards := make([]Card, 0, DeckSize-StartLayoutSize)

	for j := range layout {
		cardIndex := assists.GetRandomInteger(0, len(cardIndexes)-1)
		layout[j] = Deck[cardIndexes[cardIndex]]
		cardIndexes = assists.RemoveElement(cardIndexes, cardIndex)
	}

	sort.Slice(cardIndexes, func(i, j int) bool {
		return cardIndexes[i] < cardIndexes[j]
	})

	for _, id := range cardIndexes {
		freeCards = append(freeCards, Deck[id])
	}

	return layout, freeCards
}

func generateCards() []Card {
	cards := make([]Card, 0, DeckSize)

	for i := 0; i < len(numbers); i++ {
		for j := 0; j < len(shapes); j++ {
			for k := 0; k < len(shadings); k++ {
				for m := 0; m < len(colors); m++ {
					cards = append(cards, Card{
						Id:      len(cards),
						Number:  numbers[i],
						Shape:   shapes[j],
						Shading: shadings[k],
						Color:   colors[m],
					})
				}
			}
		}
	}

	return cards
}
