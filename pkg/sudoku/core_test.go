package sudoku

import (
	"testing"
)

var testCore *Core

func TestGenerate(t *testing.T) {
	testCore = NewCore()
	// before generation
	if testCore.Grid[0][0] != 0 {
		t.Errorf("NewCore() FAILED. Should return empty Core.Grid filled by zero values")
	} else {
		t.Logf("NewCore() PASSED")
	}

	testCore.Generate()
	// after generation
	if !(testCore.Grid[0][0] > 0) {
		t.Errorf("Core.Generate() FAILED. Should return Core.Grid filled by non-zero values")
	} else {
		t.Logf("Core.Generate() PASSED")
	}
}

func TestPrepare(t *testing.T) {
	// before prepare
	if testCore.Task[0][0] > 0 {
		t.Errorf("NewCore() FAILED. Should return empty Core.Task filled by zero values")
	} else {
		t.Logf("NewCore() PASSED")
	}

	// after preparation
	testCore.Prepare(50)
	isFilled := false
	for _, row := range testCore.Task {
		if isFilled {
			break
		}
		for _, ceil := range row {
			isFilled = ceil > 0
			if isFilled {
				break
			}
		}
	}
	if !isFilled {
		t.Errorf("testCore.Prepare(%d) FAILED. Should return Core.Task filled by some non-zero values, but got all zero values", 50)
	} else {
		t.Logf("testCore.Prepare(%d) PASSED", 50)
	}

	// 0% fill
	testValues := []int{0, -1, -100}

	for _, tv := range testValues {
		testCore.Prepare(tv)
		isFilled := false
		for _, row := range testCore.Task {
			if isFilled {
				break
			}
			for _, ceil := range row {
				isFilled = ceil > 0
				if isFilled {
					break
				}
			}
		}

		if isFilled {
			t.Errorf("testCore.Prepare(%d) FAILED. Should return Core.Task filled by zero values, but got non-zero value", tv)
		} else {
			t.Logf("testCore.Prepare(%d) PASSED", tv)
		}
	}

	// 100% fill
	testValues = []int{100, 101, 200}

	for _, tv := range testValues {
		testCore.Prepare(tv)
		isFilled := true

		for _, row := range testCore.Task {
			if !isFilled {
				break
			}
			for _, ceil := range row {
				isFilled = ceil > 0
				if !isFilled {
					break
				}
			}
		}

		if !isFilled {
			t.Errorf("testCore.Prepare(%d) FAILED. Should return Core.Task filled non-zero values, but got zero value", tv)
		} else {
			t.Logf("testCore.Prepare(%d) PASSED", tv)
		}
	}
}

func TestVerify(t *testing.T) {
	// check 100% filled task
	result := testCore.Verify()

	if !result.IsValid {
		t.Errorf("testCore.Verify() FAILED. Should be verified, but got error at position %+v", result.Position)
	} else {
		t.Logf("testCore.Verify() PASSED")
	}

	// check with error
	testCore.Task[0][0] = 1
	testCore.Task[0][1] = 1
	result = testCore.Verify()

	if result.IsValid {
		t.Errorf("testCore.Verify() FAILED. Should be an error, but got verified")
	} else {
		t.Logf("testCore.Verify() PASSED")
	}

	testCore.Prepare(100)
	testCore.Task[0][0] = 1
	testCore.Task[1][0] = 1
	result = testCore.Verify()

	if result.IsValid {
		t.Errorf("testCore.Verify() FAILED. Should be an error, but got verified")
	} else {
		t.Logf("testCore.Verify() PASSED")
	}
}
