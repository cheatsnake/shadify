package schulte

import (
	"math"
	"testing"
)

func TestGenerateNumeric(t *testing.T) {
	testValues := []int{0, 1, 10, 100, 101, 1000, -10, -101, -1000}

	for _, tv := range testValues {
		result := GenerateNumeric(tv)
		expectedSize := int(math.Abs(float64(tv)))

		if tv > 100 || tv < -100 {
			expectedSize = 100
		}

		if len(result.Grid) != expectedSize {
			t.Errorf(
				"GenerateNumeric(%d) FAILED. Expected result.Grid with length = %d, but got %d.",
				tv, expectedSize, len(result.Grid))
		} else {
			t.Logf("GenerateNumeric(%d) PASSED.", tv)
		}
	}
}

func TestGenerateAlphabetic(t *testing.T) {
	for i := 0; i < 5; i++ {
		result := GenerateAlphabetic()

		if len(result.Grid) != 5 {
			t.Errorf(
				"GenerateAlphabetic() FAILED. Expected result.Grid with length = %d, but got %d.",
				5, len(result.Grid))
		} else {
			t.Logf("GenerateAlphabetic() PASSED.")
		}
	}
}
