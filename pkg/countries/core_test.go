package countries

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
)

func init() {
	countriesJson, _ := os.ReadFile(filepath.Join("..", "..", "data", "countries.json"))
	json.Unmarshal((countriesJson), &countriesDB)
}

func TestGetCapitalQuiz(t *testing.T) {
	validVariants := []int{2, 4, 6}
	for _, vv := range validVariants {
		result, _ := GetCapitalQuiz(vv, 1)
		if len(result[0].Variants) != vv {
			t.Errorf("GetCapitalQuiz(%d, 1) FAILED. Expected result.Variants with length = %d, but got %d",
				vv, vv, len(result[0].Variants))
		} else {
			t.Logf("GetCapitalQuiz(%d, 1) PASSED.", vv)
		}
	}

	invalidVariants := []int{-2, 0, 1, 7, 10}
	for _, iv := range invalidVariants {
		_, err := GetCapitalQuiz(iv, 1)
		if err == nil {
			t.Errorf("GetCapitalQuiz(%d, 1) FAILED. Expected an error", iv)
		} else {
			t.Logf("GetCapitalQuiz(%d, 1) PASSED.", iv)
		}
	}

	validAmounts := []int{0, 1, 5, 10, 20}
	for _, va := range validAmounts {
		result, _ := GetCapitalQuiz(2, va)
		if len(result) != va {
			t.Errorf("GetCapitalQuiz(2, %d) FAILED. Expected result with length %d, but got %d", va, va, len(result))
		} else {
			t.Logf("GetCapitalQuiz(2, %d) PASSED.", va)
		}
	}

	invalidAmounts := []int{-1, -5, 21, 25}
	for _, ia := range invalidAmounts {
		_, err := GetCapitalQuiz(2, ia)
		if err == nil {
			t.Errorf("GetCapitalQuiz(2, %d) FAILED. Expected an error", ia)
		} else {
			t.Logf("GetCapitalQuiz(2, %d) PASSED.", ia)
		}
	}
}

func TestGetCountryQuiz(t *testing.T) {
	validVariants := []int{2, 4, 6}
	for _, vv := range validVariants {
		result, _ := GetCountryQuiz(vv, 1)
		if len(result[0].Variants) != vv {
			t.Errorf("GetCountryQuiz(%d) FAILED. Expected result.Variants with length = %d, but got %d",
				vv, vv, len(result[0].Variants))
		} else {
			t.Logf("GetCountryQuiz(%d) PASSED.", vv)
		}
	}

	invalidVariants := []int{-2, 0, 1, 7, 10}
	for _, iv := range invalidVariants {
		_, err := GetCountryQuiz(iv, 1)
		if err == nil {
			t.Errorf("GetCountryQuiz(%d) FAILED. Expected an error", iv)
		} else {
			t.Logf("GetCountryQuiz(%d) PASSED.", iv)
		}
	}

	validAmounts := []int{0, 1, 5, 10, 20}
	for _, va := range validAmounts {
		result, _ := GetCountryQuiz(2, va)
		if len(result) != va {
			t.Errorf("GetCountryQuiz(2, %d) FAILED. Expected result with length %d, but got %d", va, va, len(result))
		} else {
			t.Logf("GetCountryQuiz(2, %d) PASSED.", va)
		}
	}

	invalidAmounts := []int{-1, -5, 21, 25}
	for _, ia := range invalidAmounts {
		_, err := GetCountryQuiz(2, ia)
		if err == nil {
			t.Errorf("GetCountryQuiz(2, %d) FAILED. Expected an error", ia)
		} else {
			t.Logf("GetCountryQuiz(2, %d) PASSED.", ia)
		}
	}
}
