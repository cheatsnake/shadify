package set

import (
	"testing"
)

func TestNewCore(t *testing.T) {
	for i := 0; i < 3; i++ {
		sc := NewCore()
		if len(sc.FreeCards) != DeckSize {
			t.Errorf("NewCore() FAILED. Should return Core.FreeCards with length = %d, but got %d",
				DeckSize, len(sc.FreeCards))
		} else {
			t.Logf("NewCore() PASSED")
		}
	}
}

func TestStart(t *testing.T) {
	for i := 0; i < 3; i++ {
		sc := NewCore()
		sc.Start()
		if len(sc.Layout) != StartLayoutSize {
			t.Errorf("Core.Start() FAILED. Should return Core.Layout with length = %d, but got %d",
				StartLayoutSize, len(sc.Layout))
		} else if len(sc.FreeCards) != DeckSize-StartLayoutSize {
			t.Errorf("Core.Start() FAILED. Should return Core.FreeCards with length = %d, but got %d",
				DeckSize-StartLayoutSize, len(sc.FreeCards))
		} else if len(sc.WonCards) != 0 || len(sc.PossibleSets) != 0 {
			t.Errorf("Core.Start() FAILED. Should return empty Core.WonCards and Core.PossibleSets, but got lengths %d and %d",
				len(sc.WonCards), len(sc.PossibleSets))
		} else {
			t.Logf("Core.Start() PASSED")
		}
	}
}

func TestRemoveSet(t *testing.T) {
	sc := NewCore()
	sc.Start()
	sc.AddCards()
	sc.AddCards()
	sc.AddCards()
	sc.FindSets()

	err := sc.RemoveSet([]int{sc.PossibleSets[0][0].Id, sc.PossibleSets[0][1].Id, sc.PossibleSets[0][2].Id})
	if err != nil {
		t.Errorf("Core.RemoveSet() FAILED. Should return remove set, but got error: %s", err.Error())
	} else {
		t.Logf("Core.RemoveSet() PASSED")
	}

	if len(sc.WonCards) != SetSize {
		t.Errorf("Core.RemoveSet() FAILED. Should return increase Core.WonCards by %d cards, but got: %d",
			SetSize, len(sc.WonCards))
	} else {
		t.Logf("Core.RemoveSet() PASSED")
	}

	err = sc.RemoveSet([]int{sc.FreeCards[0].Id, sc.FreeCards[1].Id, sc.FreeCards[2].Id})
	if err == nil {
		t.Errorf("Core.RemoveSet() FAILED. Should return error")
	} else {
		t.Logf("Core.RemoveSet() PASSED")
	}

	isValidSet := validateSet([]int{sc.Layout[0].Id, sc.Layout[1].Id, sc.Layout[2].Id})
	for isValidSet {
		sc.Start()
		isValidSet = validateSet([]int{sc.Layout[0].Id, sc.Layout[1].Id, sc.Layout[2].Id})
	}

	err = sc.RemoveSet([]int{sc.Layout[0].Id, sc.Layout[1].Id, sc.Layout[2].Id})
	if err.Error() != setIsNotValid {
		t.Errorf("Core.RemoveSet() FAILED. Should return error: %s", setIsNotValid)
	} else {
		t.Logf("Core.RemoveSet() PASSED")
	}
}

func TestAddCards(t *testing.T) {
	sc := NewCore()
	sc.Start()

	for i := 0; i < 3; i++ {
		err := sc.AddCards()
		if err != nil {
			t.Errorf("Core.AddCards() FAILED. Should be passed, but got error: %s", err.Error())
		} else {
			t.Logf("Core.AddCards() PASSED")
		}
	}

	err := sc.AddCards()
	if err.Error() != peakMaxLayoutSize {
		t.Errorf("Core.AddCards() FAILED. Should return error: %s", peakMaxLayoutSize)
	} else {
		t.Logf("Core.AddCards() PASSED")
	}
}
