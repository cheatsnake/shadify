package memory

import "testing"

func TestLettersPool(t *testing.T) {
	t.Run("valid pair sizes", func(t *testing.T) {
		values := [][]int{
			{2, 2}, {64, 2}, {104, 2},
			{6, 3}, {60, 3}, {156, 3},
			{4, 4}, {80, 4}, {208, 4},
		}

		for _, val := range values {
			letters, err := lettersPool(val[0], val[1])
			if err != nil {
				t.Errorf("should be valid, but got error: %v", err.Error())
			}
			if len(letters) != val[0] {
				t.Errorf("total number of letters should be equal to %d, but got %d", val[0], len(letters))
			}
		}
	})
}
