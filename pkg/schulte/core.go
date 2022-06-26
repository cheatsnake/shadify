package schulte

import "github.com/cheatsnake/shadify/pkg/assists"

func NewCore() *Core {
	return &Core{}
}

func (sc *Core) GenerateNumeric(size int) [][]int {
	values := assists.GetNumbers(size * size, 1)
	return Generate(values)
}

func (sc *Core) GenerateAlphabetic() [][]string {
	values := make([]string, 25)
	copy(values, alphabet[0:25])
	return Generate(values)
}