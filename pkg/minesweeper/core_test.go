package minesweeper

import "testing"

type TestValue struct {
	sp          string
	w, h, mines int
}

func TestGenerateSuccess(t *testing.T) {
	testValues := []TestValue{
		{"1-1", 9, 9, 12},
		{"1-1", 16, 16, 64},
		{"1-1", 20, 12, 60},
		{"1-1", 4, 4, 4},
	}

	for _, tv := range testValues {
		result, _ := Generate(tv.sp, tv.w, tv.h, tv.mines)
		if len(result.Board) != tv.h || len(result.Board[0]) != tv.w || result.Board[0][0] != "o" {
			t.Errorf(
				"Generate(%s, %d, %d, %d) FAILED. Expected w=%d, h=%d, sp='o', but got w=%d, h=%d, sp=%+v\n",
				tv.sp, tv.w, tv.h, tv.mines, tv.w, tv.h, len(result.Board), len(result.Board[0]), result.Board[0][0])
		} else {
			t.Logf(
				"Generate(%s, %d, %d, %d) PASSED\n",
				tv.sp, tv.w, tv.h, tv.mines)
		}
	}
}

func TestGenerateFail(t *testing.T) {
	testValues := []TestValue{
		{"1-1", 1, 1, 1},
		{"1-1", 32, 32, 32},
		{"1-1", 10, 10, 26},
		{"13-13", 12, 12, 12},
	}

	for _, tv := range testValues {
		result, err := Generate(tv.sp, tv.w, tv.h, tv.mines)
		if err == nil {
			t.Errorf(
				"Generate(%s, %d, %d, %d) FAILED. Expected error, but got %+v",
				tv.sp, tv.w, tv.h, tv.mines, result)
		} else {
			t.Logf(
				"Generate(%s, %d, %d, %d) -> '%+v' PASSED\n",
				tv.sp, tv.w, tv.h, tv.mines, err)
		}
	}

}
