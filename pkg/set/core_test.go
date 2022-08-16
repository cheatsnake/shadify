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

type testState struct {
	State  string
	Expect string
}

func TestFindSets(t *testing.T) {
	sc := NewCore()
	testStates := []testState{
		{State: "12-79-77-49-1-80-13-66-16-17-10-30", Expect: "found 4 sets"},
		{State: "69-42-54-43-59-20-38-63-46-61-4-50", Expect: "found 1 set"},
		{State: "14-74-48-45-32-22-12-46-26-71-24-44", Expect: setsNotFound},
		{State: "11-50-60@0-1-2-3-4-5-6-7-8-9-10-12-13-14-15-16-17-18-19-20-21-22-23-24-25-26-27-28-29-30-31-32-33-34-35-36-37-38-39-40-41-42-43-44-45-46-47-48-49-51-52-53-54-55-56-57-58-59-61-62-63-64-65-66-67-68-69-70-71-72-73-74-75-76-77-78-79-80", Expect: gameOver}}

	for _, ts := range testStates {
		sc.LoadState(ts.State)
		msg := sc.FindSets()
		if msg != ts.Expect {
			t.Errorf("Core.FindSets() FAILED. Should return message: %s, but got: %s", ts.Expect, msg)
		} else {
			t.Logf("Core.FindSets() PASSED")
		}
	}
}

func TestLoadState(t *testing.T) {
	sc := NewCore()
	testStates := []string{
		"a-a-a",
		"0-1-2",
		"0-1-2-3-4-5-6-7-8-9-10-11-12",
		"0-1-2-3-4-5-6-7-8-9-10-81",
		"0-1-2-3-4-5-6-7-8-9-10-11@0-1-2",
		"0-1-2-3-4-5-6-7-8-9-10-11@12-13",
		"0-1-2-3-4-5-6-7-8-9-10-11@12-13-81",
	}

	for _, ts := range testStates {
		result := sc.LoadState(ts)
		if result == nil {
			t.Errorf("Core.AddCards() FAILED. Should return error")
		} else {
			t.Logf("Core.LoadState() PASSED")
		}
	}
}
