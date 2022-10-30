package countries

import (
	"encoding/json"
	"errors"
	"log"
	"math"
	"os"
	"path/filepath"
)

var countriesDB []CountryData

func init() {
	countriesJson, err := os.ReadFile(filepath.Join("data", "countries.json"))
	if err != nil {
		log.Printf("error opening file: %v \n", err)
	}

	err = json.Unmarshal((countriesJson), &countriesDB)
	if err != nil {
		log.Printf("error parsing JSON: %v \n", err)
	}
}

// Get a quiz in which you have to guess the capital of the specified country
func GetCapitalQuiz(variants, amount int) ([]CapitalQuiz, error) {
	if variants > maxVariantsCount {
		return []CapitalQuiz{}, errors.New(variantsCountToBig)
	}

	if variants < minVariantsCount {
		return []CapitalQuiz{}, errors.New(variantsCountToSmall)
	}

	if amount > maxQuizAmount {
		return []CapitalQuiz{}, errors.New(quizAmountToBig)
	}

	if amount < 0 {
		return []CapitalQuiz{}, errors.New(quizAmountToSmall)
	}

	result := generateCapitalQuizzes(variants, int(math.Abs(float64(amount))))
	return result, nil
}

// Get a quiz in which you have to guess the name of the country from the image of its flag
func GetCountryQuiz(variants, amount int) ([]CountryQuiz, error) {
	if variants > maxVariantsCount {
		return []CountryQuiz{}, errors.New(variantsCountToBig)
	}

	if variants < minVariantsCount {
		return []CountryQuiz{}, errors.New(variantsCountToSmall)
	}

	if amount > maxQuizAmount {
		return []CountryQuiz{}, errors.New(quizAmountToBig)
	}

	if amount < 0 {
		return []CountryQuiz{}, errors.New(quizAmountToSmall)
	}

	result := generateCountryQuizzes(variants, int(math.Abs(float64(amount))))
	return result, nil
}
