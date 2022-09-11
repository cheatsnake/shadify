package countries

import (
	"encoding/json"
	"io/ioutil"
	"path/filepath"
	"testing"
)

func init() {
	countriesJson, _ := ioutil.ReadFile(filepath.Join("..", "..", "data", "countries.json"))
	json.Unmarshal((countriesJson), &countriesDB)
}

func TestGetCapitalQuiz(t *testing.T) {
	validValues := []int{2, 4, 6}
	for _, vv := range validValues {
		result, _ := GetCapitalQuiz(vv, 1)
		if len(result[0].Variants) != vv {
			t.Errorf("GetCapitalQuiz(%d) FAILED. Expected result.Variants with length = %d, but got %d",
				vv, vv, len(result[0].Variants))
		} else {
			t.Logf("GetCapitalQuiz(%d) PASSED.", vv)
		}
	}

	invalidValues := []int{-2, 0, 1, 7, 10}
	for _, iv := range invalidValues {
		_, err := GetCapitalQuiz(iv, 1)
		if err == nil {
			t.Errorf("GetCapitalQuiz(%d) FAILED. Expected error", iv)
		} else {
			t.Logf("GetCapitalQuiz(%d) -> %s PASSED.", iv, err.Error())
		}
	}
}

func TestGetCountryQuiz(t *testing.T) {
	validValues := []int{2, 4, 6}
	for _, vv := range validValues {
		result, _ := GetCountryQuiz(vv, 1)
		if len(result[0].Variants) != vv {
			t.Errorf("GetCountryQuiz(%d) FAILED. Expected result.Variants with length = %d, but got %d",
				vv, vv, len(result[0].Variants))
		} else {
			t.Logf("GetCountryQuiz(%d) PASSED.", vv)
		}
	}

	invalidValues := []int{-2, 0, 1, 7, 10}
	for _, iv := range invalidValues {
		_, err := GetCountryQuiz(iv, 1)
		if err == nil {
			t.Errorf("GetCountryQuiz(%d) FAILED. Expected error", iv)
		} else {
			t.Logf("GetCountryQuiz(%d) -> %s PASSED.", iv, err.Error())
		}
	}
}
