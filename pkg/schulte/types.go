package schulte

// Core struct of Schulte module
type Core[T int | string] struct {
	Grid [][]T `json:"grid"` // Schulte grid with randomly placed numbers or letters
}
