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
var errPairSizeNotAllowed = errors.New("pair size not allowed")
var errTooManyLetters = errors.New("total number of letters must not exceed 52 multiplied by the size of the pair")
var errIncorrectLettersAmount = errors.New("total number of letters must be even to the size of the pair")
