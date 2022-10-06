package camp

import (
	"testing"
)

func TestGenerate(t *testing.T) {
	testValidValues := [][]int{{5, 5}, {7, 7}, {15, 15}, {5, 15}, {15, 5}}

	for _, tvv := range testValidValues {
		_, err := Generate(tvv[0], tvv[1])

		if err != nil {
			t.Errorf("Generate(%d, %d) FAILED. Expected error == nil, but got %v", tvv[0], tvv[1], err.Error())
		} else {
			t.Logf("Generate(%d, %d) PASSED", tvv[0], tvv[1])
		}
	}

	testInvalidValues := [][]int{{0, 0}, {4, 4}, {16, 16}, {10, 20}, {20, 10}}

	for _, tiv := range testInvalidValues {
		_, err := Generate(tiv[0], tiv[1])

		if err == nil {
			t.Errorf("Generate (%d, %d) FAILED. Expected error but got nil", tiv[0], tiv[1])
		} else {
			t.Logf("Generate (%d, %d) PASSED", tiv[0], tiv[1])
		}
	}
}
