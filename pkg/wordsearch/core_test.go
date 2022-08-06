package wordsearch

import (
	"path/filepath"
	"testing"

	"github.com/cheatsnake/shadify/pkg/assists"
)

func TestGenerateSuccess(t *testing.T) {
	wordsDB, _ = assists.ReadFileLineByLine(filepath.Join("..", "..", "data", "nouns.txt"))
	testValues := [][]int{{8, 8}, {5, 5}, {16, 16}, {20, 12}, {12, 20}, {5, 20}, {20, 5}}

	for _, tv := range testValues {
		result, _ := Generate(tv[0], tv[1])

		if result.Width != tv[0] ||
			result.Height != tv[1] ||
			len(result.Grid)*len(result.Grid[0]) != tv[0]*tv[1] {
			t.Errorf(
				"Generate(%d, %d) FAILED. Expected result.Width = %d, result.Height = %d and total ceils = %d, but got %d, %d, %d",
				tv[0], tv[1], tv[0], tv[1], tv[0]*tv[1], result.Width, result.Height, len(result.Grid)*len(result.Grid[0]))
		} else {
			t.Logf("Generate(%d, %d) PASSED", tv[0], tv[1])
		}
	}
}

func TestGenerateFail(t *testing.T) {
	testValues := [][]int{{4, 8}, {21, 8}, {8, 4}, {8, 21}, {17, 16}}

	for _, tv := range testValues {
		_, err := Generate(tv[0], tv[1])

		if err == nil {
			t.Errorf("Generate(%d, %d) FAILED. Expected error != nil, but got nil", tv[0], tv[1])
		} else {
			t.Logf("Generate(%d, %d) -> '%+v', PASSED", tv[0], tv[1], err)
		}
	}
}
