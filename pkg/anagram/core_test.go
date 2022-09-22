package anagram

import (
	"path/filepath"
	"testing"

	"github.com/cheatsnake/shadify/internal/helpers"
)

func TestGenerate(t *testing.T) {
	wordsDB, _ = helpers.ReadFileLineByLine(filepath.Join("..", "..", "data", "nouns.txt"))
	wordsDBSize = len(wordsDB)
	result := Generate()

	for i := 0; i < 10; i++ {
		if !(len(result.Task) >= 9 && len(result.Words) > 0) {
			t.Errorf(
				"Generate() FAILED. Expected result.Task with length >= 9, but got result.Task length = %d and result.Words length = %d.",
				len(result.Task), len(result.Words))
		} else {
			t.Logf("Generate() PASSED.")
		}
	}
}
