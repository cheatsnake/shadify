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
