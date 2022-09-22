package set

import (
	"errors"
	"strconv"
	"strings"

	"github.com/cheatsnake/shadify/internal/helpers"
)

func updateState(layout []Card, wonCards []Card) string {
	layoutId := make([]string, 0, DeckSize)
	wonCardsId := make([]string, 0, DeckSize)

	for _, card := range layout {
		layoutId = append(layoutId, strconv.Itoa(card.Id))
	}

	if len(wonCards) < 1 {
		return strings.Join(layoutId[:], idSeparator)
	}

	for _, card := range wonCards {
		wonCardsId = append(wonCardsId, strconv.Itoa(card.Id))
	}

	return strings.Join(layoutId[:], idSeparator) + stateSeparator + strings.Join(wonCardsId[:], idSeparator)
}

func loadState(state string) ([]Card, []Card, []Card, error) {
	layoutId, wonCardsId, err := parseStateString(state)
	if err != nil {
		return nil, nil, nil, err
	}

	if len(layoutId)+len(wonCardsId) < StartLayoutSize {
		return nil, nil, nil, errors.New(notEnoughCards)
	}

	if len(layoutId) > MaxLayerSize {
		return nil, nil, nil, errors.New(layoutIsTooBig)
	}

	if len(layoutId)%SetSize != 0 || len(wonCardsId)%SetSize != 0 {
		return nil, nil, nil, errors.New(incorrectCardsAmount)
	}

	if len(layoutId)+len(wonCardsId) > DeckSize {
		return nil, nil, nil, errors.New(tooMuchCards)
	}

	layout := make([]Card, 0, DeckSize)
	freeCards := make([]Card, 0, DeckSize-StartLayoutSize)
	wonCards := make([]Card, 0, DeckSize)

	for j := range layoutId {
		layout = append(layout, Deck[layoutId[j]])
	}

	for _, id := range wonCardsId {
		wonCards = append(wonCards, Deck[id])
	}

	usedCardsId := append(layoutId, wonCardsId...)

	for i := range Deck {
		if !helpers.SliceContains(usedCardsId, i) {
			freeCards = append(freeCards, Deck[i])
		}
	}

	return layout, freeCards, wonCards, nil
}

func parseStateString(state string) ([]int, []int, error) {
	splitedState := strings.Split(state, stateSeparator)
	layoutIdStr := helpers.RemoveDuplicateStrings(strings.Split(splitedState[0], "-"))
	layoutId, err := helpers.SliceStringToInt(layoutIdStr)
	if err != nil {
		return nil, nil, err
	}

	for _, id := range layoutId {
		if id >= DeckSize {
			return nil, nil, errors.New("entered not exist id: " + strconv.Itoa(id))
		}
	}

	if len(splitedState) < 2 {
		return layoutId, nil, nil
	}

	wonCardsIdStr := helpers.RemoveDuplicateStrings(strings.Split(splitedState[1], "-"))
	wonCardsId, err := helpers.SliceStringToInt(wonCardsIdStr)
	if err != nil {
		return nil, nil, err
	}

	for _, id := range wonCardsId {
		if id >= DeckSize {
			return nil, nil, errors.New("entered not exist id: " + strconv.Itoa(id))
		}
		if helpers.SliceContains(layoutId, id) {
			return nil, nil, errors.New("layout must not contain won card id: " + strconv.Itoa(id))
		}
	}

	return layoutId, wonCardsId, nil
}
