package memory

import (
	"errors"
	"testing"
)

func TestGenerate(t *testing.T) {
	t.Run("valid params", func(t *testing.T) {
		values := [][]int{
			{2, 2, 2, 0},
			{2, 2, 4, 1},
			{6, 4, 3, 0},
			{9, 6, 2, 1},
			{13, 8, 2, 0},
			{12, 13, 3, 1},
			{16, 13, 4, 0},
		}

		for _, v := range values {
			_, err := Generate(v[0], v[1], v[2], v[3] == 1)
			if err != nil {
				t.Errorf("should be fine, but got error: %s, with this params: %v", err.Error(), v)
			}
		}
	})

	t.Run("incorrect pair size", func(t *testing.T) {
		values := [][]int{
			{2, 2, 1, 0},
			{2, 2, 5, 1},
			{6, 4, -2, 0},
			{9, 6, 10, 1},
			{13, 8, 0, 0},
			{12, 13, -3, 1},
			{16, 13, 12, 0},
		}

		for _, v := range values {
			_, err := Generate(v[0], v[1], v[2], v[3] == 1)
			if err == nil {
				t.Errorf("should be error, but got fine with this params: %v", v)
			}
			if !errors.Is(err, errPairSizeNotAllowed) {
				t.Errorf("should be %s, but got %s", errPairSizeNotAllowed.Error(), err.Error())
			}
		}
	})

	t.Run("nagative grid size", func(t *testing.T) {
		values := [][]int{
			{-2, 4, 2, 0},
			{4, -2, 3, 1},
			{-10, -4, 4, 0},
		}

		for _, v := range values {
			_, err := Generate(v[0], v[1], v[2], v[3] == 1)
			if err == nil {
				t.Errorf("should be error, but got fine with this params: %v", v)
			}
			if !errors.Is(err, errNegativeNumbers) {
				t.Errorf("should be %s, but got %s", errNegativeNumbers.Error(), err.Error())
			}
		}
	})

	t.Run("too big grid size", func(t *testing.T) {
		values := [][]int{
			{9, 13, 2, 0},
			{13, 9, 2, 1},
			{12, 14, 3, 0},
			{14, 12, 3, 1},
			{16, 14, 4, 0},
			{14, 16, 4, 1},
		}

		for _, v := range values {
			_, err := Generate(v[0], v[1], v[2], v[3] == 1)
			if err == nil {
				t.Errorf("should be error, but got fine with this params: %v", v)
			}
			if !errors.Is(err, errTooBigGrid) {
				t.Errorf("should be %s, but got %s", errTooBigGrid.Error(), err.Error())
			}
		}
	})

	t.Run("incorrect grid size", func(t *testing.T) {
		values := [][]int{
			{3, 5, 2, 0},
			{10, 4, 3, 1},
			{5, 5, 4, 0},
		}

		for _, v := range values {
			_, err := Generate(v[0], v[1], v[2], v[3] == 1)
			if err == nil {
				t.Errorf("should be error, but got fine with this params: %v", v)
			}
			if !errors.Is(err, errIncorrectGridSize) {
				t.Errorf("should be %s, but got %s", errIncorrectGridSize.Error(), err.Error())
			}
		}
	})
}
