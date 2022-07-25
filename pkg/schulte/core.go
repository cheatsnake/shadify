package schulte

import "github.com/cheatsnake/shadify/pkg/assists"

// Generate a Schulte numerical grid with a given size
func GenerateNumeric(size int) Core[int] {
	values := assists.GetNumbers(size*size, 1)
	grid := generateGrid(values)

	return Core[int]{Grid: grid}
}

// Generate a Schulte alphabetic grid with fixed 5x5 size
func GenerateAlphabetic() Core[string] {
	values := make([]string, 25)
	copy(values, alphabet[0:25])
	grid := generateGrid(values)

	return Core[string]{Grid: grid}
}
