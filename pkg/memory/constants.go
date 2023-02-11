package memory

import "errors"

const (
	_lowerCaseLetters string = "abcdefghijklmnopqrstuvwxyz"
	_upperCaseLetters string = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

var _pairSizes = map[int]int{
	2: 2,
	3: 3,
	4: 4,
}

var errNegativeNumbers = errors.New("negative numbers are not allowed")
var errPairSizeNotAllowed = errors.New("given pair size not allowed")
var errTooBigGrid = errors.New("total number of cells in the grid must not exceed 52 multiplied by the pair size")
var errIncorrectGridSize = errors.New("total number of cells in the grid must be even to the size of the pair")
