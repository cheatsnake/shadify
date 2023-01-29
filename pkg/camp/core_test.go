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

type verifyArgs struct {
	RowTents    []int
	ColumnTents []int
	Solution    [][]int
}

func TestVerify(t *testing.T) {
	testValidValue := verifyArgs{[]int{1, 1, 1, 1, 1}, []int{1, 1, 0, 3, 0}, [][]int{{1, 0, 0, 2, 0}, {2, 0, 0, 1, 0}, {0, 0, 0, 2, 1}, {0, 2, 0, 0, 0}, {0, 1, 0, 2, 1}}}
	result, err := Verify(testValidValue.Solution, testValidValue.RowTents, testValidValue.ColumnTents)

	if err != nil || result.IsValid {
		t.Errorf("Verify(...) with valid args FAILED. Expected error != nil & result.IsValid = false, but got %v, %v", err, result.IsValid)
	} else {
		t.Logf("Verify(...) PASSED")
	}

	testInvalidValue := verifyArgs{[]int{1, 1, 1, 1, 1}, []int{1, 1, 0, 3, 0}, [][]int{{1, 2, 2, 2, 0}, {2, 0, 0, 1, 0}, {0, 0, 0, 2, 1}, {0, 2, 0, 0, 0}, {0, 1, 0, 2, 1}}}

	result, err = Verify(testInvalidValue.Solution, testInvalidValue.RowTents, testInvalidValue.ColumnTents)

	if err != nil || !result.IsValid {
		t.Errorf("Verify(...) with invalid args FAILED. Expected error = nil & result.IsValid = true, but got %v, %v", err, result.IsValid)
	} else {
		t.Logf("Verify(...) PASSED")
	}

	testIncorrectValues := []verifyArgs{
		{[]int{1, 1, 1, 1},
			[]int{1, 1, 0, 3, 0},
			[][]int{{1, 2, 2, 2, 0}, {2, 0, 0, 1, 0}, {0, 0, 0, 2, 1}, {0, 2, 0, 0, 0}, {0, 1, 0, 2, 1}}},
		{[]int{1, 1, 1, 1, 1},
			[]int{1, 1, 0, 3},
			[][]int{{1, 2, 2, 2, 0}, {2, 0, 0, 1, 0}, {0, 0, 0, 2, 1}, {0, 2, 0, 0, 0}, {0, 1, 0, 2, 1}}},
		{[]int{0},
			[]int{0},
			[][]int{{1}, {1}, {1}, {1}, {1}, {1}, {1}, {1}, {1}, {1}, {1}, {1}, {1}, {1}, {1}, {1}}},
		{[]int{0},
			[]int{0},
			[][]int{{1}}},
	}

	for _, tiv := range testIncorrectValues {
		result, err = Verify(tiv.Solution, tiv.RowTents, tiv.ColumnTents)

		if err == nil {
			t.Errorf("Verify(...) with incorrect args FAILED. Expected error != nil, but got nil")
		} else {
			t.Logf("Verify(...) PASSED")
		}
	}
}
