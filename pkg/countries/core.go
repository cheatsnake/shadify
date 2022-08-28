package countries

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"path/filepath"
)

var countriesDB []CountryData

func init() {
	countriesJson, err := ioutil.ReadFile(filepath.Join("data", "countries.json"))
	if err != nil {
		log.Printf("error opening file: %v \n", err)
	}

	err = json.Unmarshal((countriesJson), &countriesDB)
	if err != nil {
		log.Printf("error parsing JSON: %v \n", err)
	}
}

// Get a quiz in which you have to guess the capital of the specified country
func GetCapitalQuiz(variants int) (CapitalQuiz, error) {
	if variants > maxVariantsCount {
		return CapitalQuiz{}, errors.New(variantsCountToBig)
	}

	if variants < minVariantsCount {
		return CapitalQuiz{}, errors.New(variantsCountToLow)
	}

	result := generateCapitalQuiz(variants)
	return result, nil
}

// Get a quiz in which you have to guess the name of the country from the image of its flag
func GetCountryQuiz(variants int) (CountryQuiz, error) {
	if variants > maxVariantsCount {
		return CountryQuiz{}, errors.New(variantsCountToBig)
	}

	if variants < minVariantsCount {
		return CountryQuiz{}, errors.New(variantsCountToLow)
	}

	result := generateCountryQuiz(variants)
	return result, nil
}
