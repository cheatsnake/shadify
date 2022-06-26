package assists

import (
	"errors"
	"math/rand"
	"strconv"
)

func GetRandomInteger(min, max int) int {
	if min > max {
		max = min
	}
	result := rand.Intn(max + 1 - min) + min
	return result
}

func PadStart(str string, count int, padStr string) string {
	result := ""
	for i := 0; i < count; i++ {
		result += padStr
	}
	result += str
	return result
}

func IndexOf[T int|string](element T, data []T) int {
	for i, value := range data {
		if element == value {
			return i
		}
	}
	return -1 
 }

 func GetRandomBool(chance int) bool {
	return rand.Intn(100) < chance
 }

 func SliceContains[T int | string](a []T, x T) bool {
    for _, n := range a {
        if x == n {
            return true
        }
    }
    return false
}

func RemoveElement[T any](s []T, i int) []T {
    s[i] = s[len(s)-1]
    return s[:len(s)-1]
}

func GetNumbers(amount int, startFrom int) []int {
	numbers := make([]int, amount)

	for i := 0; i < amount; i++ {
		numbers[i] = i + startFrom
	}

	return numbers
}

func SliceStringToInt(src []string) ([]int, error) {
	dst := make([]int, len(src))
	for i, item := range(src) {
		item, err := strconv.Atoi(item)
		if err != nil {
			return nil, errors.New("failed to convert string to int")
		}
		dst[i] = item
	}
	return dst, nil
}

func RemoveDuplicateStrings(strSlice []string) []string {
    allKeys := make(map[string]bool)
    list := []string{}
    for _, item := range strSlice {
        if _, value := allKeys[item]; !value {
            allKeys[item] = true
            list = append(list, item)
        }
    }
    return list
}