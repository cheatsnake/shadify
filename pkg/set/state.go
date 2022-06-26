package set

import (
	"errors"
	"strconv"
	"strings"

	"github.com/cheatsnake/shadify/pkg/assists"
)

func updateState(layout []Card, wonCards[]Card) string {
	layoutId := make([]string, 0, deckSize)
	wonCardsId := make([]string, 0, deckSize)

	for _, card := range(layout) {
		layoutId = append(layoutId, strconv.Itoa(card.Id))
	}

	if (len(wonCards) < 1) {
		return strings.Join(layoutId[:], idSeparator)
	}

	for _, card := range(wonCards) {
		wonCardsId = append(wonCardsId, strconv.Itoa(card.Id))
	}

	return strings.Join(layoutId[:], idSeparator) + stateSeparator + strings.Join(wonCardsId[:], idSeparator)
}

func loadState(state string) ([]Card, []Card, []Card, error) {
	layoutId, wonCardsId, err := parseStateString(state)
	if err != nil {
		return nil, nil, nil, err 
	}

	if len(layoutId) + len(wonCardsId) < startLayoutSize {
		return nil, nil, nil, errors.New("the size of the layout and the number of won cards in total must have at least 12 cards")
	}

	if len(layoutId) > maxLayerSize {
		return nil, nil, nil, errors.New("layout is too big (the maximum size of the layout is 21 cards)")
	}

	if len(layoutId) % setSize != 0 || len(wonCardsId) % setSize != 0 {
		return nil, nil, nil, errors.New("the layout size and the number of won cards must be a multiple of 3")
	}

	if len(layoutId) + len(wonCardsId) > deckSize {
		return nil, nil, nil, errors.New("cannot be more than 81 cards in the game")
	}

	layout := make([]Card, 0, deckSize)
	freeCards := make([]Card, 0, deckSize - startLayoutSize)
	wonCards := make([]Card, 0, deckSize)
	
	for j := range layoutId {
		layout= append(layout, Deck[layoutId[j]])
	}



	for _, id := range wonCardsId {
		wonCards = append(wonCards, Deck[id])
	}

	usedCardsId := append(layoutId, wonCardsId...)

	for i := range Deck {
		if !assists.SliceContains(usedCardsId, i) {
			freeCards = append(freeCards, Deck[i])
		}
	}

	return layout, freeCards, wonCards, nil
}

func parseStateString(state string) ([]int, []int, error) {
	splitedState := strings.Split(state, stateSeparator)
	layoutIdStr := assists.RemoveDuplicateStrings(strings.Split(splitedState[0], "-"))
	layoutId, err := assists.SliceStringToInt(layoutIdStr)
	if err != nil {
		return nil, nil, err
	} 

	for _, id := range layoutId {
		if (id >= deckSize) {
			return nil, nil, errors.New("entered not exist id: " + strconv.Itoa(id))
		}
	}

	if len(splitedState) < 2 {
		return layoutId, nil, nil
	}

	wonCardsIdStr := assists.RemoveDuplicateStrings(strings.Split(splitedState[1], "-"))
	wonCardsId, err := assists.SliceStringToInt(wonCardsIdStr)
	if err != nil {
		return nil, nil, err
	} 

	for _, id := range wonCardsId {
		if (id >= deckSize) {
			return nil, nil, errors.New("entered not exist id: " + strconv.Itoa(id))
		}
		if assists.SliceContains(layoutId, id) {
			return nil, nil, errors.New("layout must not contain won card id: " + strconv.Itoa(id))
		}
	}

	return layoutId, wonCardsId, nil
}