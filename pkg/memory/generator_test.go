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

	t.Run("invalid pair sizes", func(t *testing.T) {
		values := [][]int{
			{2, 3}, {64, 3}, {104, 3},
			{6, 4}, {60, 5}, {156, 1},
			{4, 3}, {80, 3}, {208, 3},
		}

		for _, val := range values {
			_, err := lettersPool(val[0], val[1])
			if err == nil {
				t.Error("should be invalid, but got valid")
			}
		}
	})

	t.Run("invalid total letters", func(t *testing.T) {
		values := [][]int{
			{105, 2}, {2, 200}, {157, 3}, {200, 3}, {209, 4}, {250, 4},
		}

		for _, val := range values {
			_, err := lettersPool(val[0], val[1])
			if err == nil {
				t.Error("should be invalid, but got valid")
			}
		}
	})
}
