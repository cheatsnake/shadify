package set

import (
	"errors"
	"strconv"

	"github.com/cheatsnake/shadify/pkg/assists"
)

// Create a new instance of Set Core
func NewCore() *Core {
	return &Core{FreeCards: Deck}
}

// Start a new game with 12 randomly selected cards for Core.Layout
func (sc *Core) Start() {
	sc.Layout, sc.FreeCards = startGame()
	sc.PossibleSets = [][]Card{}
	sc.WonCards = []Card{}
	sc.State = updateState(sc.Layout, sc.WonCards)
}

// Remove a set from Core.Layout by passing an array of card IDs
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
		return errors.New(setIsNotValid)
	}

	for _, id := range set {
		sc.WonCards = append(sc.WonCards, Deck[id])
		index := assists.IndexOf(id, layoutId)
		sc.Layout = assists.RemoveElement(sc.Layout, index)
		layoutId = assists.RemoveElement(layoutId, index)
	}

	if len(sc.Layout) < StartLayoutSize && len(sc.FreeCards) >= SetSize {
		sc.AddCards()
	}

	sc.PossibleSets = [][]Card{}
	sc.State = updateState(sc.Layout, sc.WonCards)

	return nil
}

// Add 3 random cards to Core.layout from Core.FreeCards
func (sc *Core) AddCards() error {
	if len(sc.Layout) >= MaxLayerSize {
		return errors.New(peakMaxLayoutSize)
	}

	if len(sc.FreeCards) < SetSize {
		return errors.New(freeCardsOver)
	}

	for i := 0; i < SetSize; i++ {
		index := assists.GetRandomInteger(0, len(sc.FreeCards) - 1)
		sc.Layout = append(sc.Layout, sc.FreeCards[index])
		sc.FreeCards = assists.RemoveElement(sc.FreeCards, index)
	}

	sc.PossibleSets = [][]Card{}
	sc.State = updateState(sc.Layout, sc.WonCards)

	return nil
}

// Find possible sets in current Core.Layout and them in Core.PossibleSets
func (sc *Core) FindSets() string {
	sc.PossibleSets = findSets(sc.Layout)	
	if len(sc.FreeCards) < SetSize && len(sc.PossibleSets) < 1 {
		return gameOver
	} else {
		if (len(sc.PossibleSets) > 1) {
			return "found " + strconv.Itoa(len(sc.PossibleSets)) + " sets"
		} else if (len(sc.PossibleSets) == 1) {
			return "found 1 set"
		} else {
			return setsNotFound
		} 
	}
}

// Load state of the game by passing state-string
func (sc *Core) LoadState(state string) error {
	layout, freeCards, wonCards, err := loadState(state)
	if err != nil {
		return err
	}

	sc.Layout = layout
	sc.FreeCards = freeCards
	sc.WonCards = wonCards
	sc.State = state
	sc.PossibleSets = [][]Card{}

	return nil
}