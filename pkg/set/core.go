package set

import (
	"errors"
	"strconv"

	"github.com/cheatsnake/shadify/pkg/assists"
)

func NewCore() *Core {
	return &Core{FreeCards: Deck}
}

func (sc *Core) Start() {
	sc.Layout, sc.FreeCards = startGame()
	sc.State = updateState(sc.Layout, sc.WonCards)
}

func (sc *Core) RemoveSet(set []int) error {
	layoutId := make([]int, len(sc.Layout))

	for i, card := range sc.Layout {
		layoutId[i] = card.Id
	}

	for _, id := range set {
		if !assists.SliceContains(layoutId, id) {
			return errors.New("the card with id:" + strconv.Itoa(id) + " was not found in the current layout")
		}
	}

	if !validateSet(set) {
		return errors.New("this combination is not a set")
	}

	for _, id := range set {
		sc.WonCards = append(sc.WonCards, Deck[id])
		index := assists.IndexOf(id, layoutId)
		sc.Layout = assists.RemoveElement(sc.Layout, index)
		layoutId = assists.RemoveElement(layoutId, index)
	}

	if len(sc.Layout) < startLayoutSize && len(sc.FreeCards) >= setSize {
		sc.AddCards()
	}

	sc.PossibleSets = [][]Card{}
	sc.State = updateState(sc.Layout, sc.WonCards)

	return nil
}

func (sc *Core) AddCards() error {
	if len(sc.Layout) > 20 {
		return errors.New("in a layout the size of more then 20 cards is guaranteed to be a set")
	}

	if len(sc.FreeCards) < setSize {
		return errors.New("there are no more free cards left")
	}

	for i := 0; i < setSize; i++ {
		index := assists.GetRandomInteger(0, len(sc.FreeCards) - 1)
		sc.Layout = append(sc.Layout, sc.FreeCards[index])
		sc.FreeCards = assists.RemoveElement(sc.FreeCards, index)
	}

	sc.PossibleSets = [][]Card{}
	sc.State = updateState(sc.Layout, sc.WonCards)

	return nil
}

func (sc *Core) FindSets() string {
	sc.PossibleSets = findSets(sc.Layout)	
	if len(sc.FreeCards) < setSize && len(sc.PossibleSets) < 1 {
		return "the game is over there are no more sets"
	} else {
		if (len(sc.PossibleSets) > 1) {
			return "found " + strconv.Itoa(len(sc.PossibleSets)) + " sets"
		} else if (len(sc.PossibleSets) == 1) {
			return "found 1 set"
		} else {
			return "sets not found"
		} 
	}
}

func (sc *Core) LoadState(state string) error {
	layout, freeCards, wonCards, err := loadState(state)
	if err != nil {
		return err
	}

	sc.Layout = layout
	sc.FreeCards = freeCards
	sc.WonCards = wonCards
	sc.State = state

	return nil
}