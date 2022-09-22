package countries

import (
	"math/rand"
	"strings"

	"github.com/cheatsnake/shadify/internal/helpers"
)

func generateCapitalQuiz(variantsCount int) CapitalQuiz {
	variants := make([]string, variantsCount)
	randomCountryIdx := rand.Intn(len(countriesDB))
	randomCountry := countriesDB[randomCountryIdx]
	flag := flagImageUrl + strings.ToLower(randomCountry.Code) + flagImageFormat
	variants[0] = randomCountry.Capital

	for i := 1; i < variantsCount; i += 0 {
		randomIdx := rand.Intn(len(countriesDB))
		capital := countriesDB[randomIdx].Capital

		if helpers.IndexOf(capital, variants) == -1 {
			variants[i] = countriesDB[randomIdx].Capital
			i++
		}
	}

	rand.Shuffle(len(variants), func(i, j int) {
		variants[i], variants[j] = variants[j], variants[i]
	})

	return CapitalQuiz{
		Country:  randomCountry.Country,
		Flag:     flag,
		Variants: variants,
		Answer:   randomCountry.Capital}
}

func generateCountryQuiz(variantsCount int) CountryQuiz {
	variants := make([]string, variantsCount)
	randomCountryIdx := rand.Intn(len(countriesDB))
	randomCountry := countriesDB[randomCountryIdx]
	flag := flagImageUrl + strings.ToLower(randomCountry.Code) + flagImageFormat
	variants[0] = randomCountry.Country

	for i := 1; i < variantsCount; i += 0 {
		randomIdx := rand.Intn(len(countriesDB))
		country := countriesDB[randomIdx].Country

		if helpers.IndexOf(country, variants) == -1 {
			variants[i] = countriesDB[randomIdx].Country
			i++
		}
	}

	rand.Shuffle(len(variants), func(i, j int) {
		variants[i], variants[j] = variants[j], variants[i]
	})

	return CountryQuiz{
		Flag:     flag,
		Variants: variants,
		Answer:   randomCountry.Country}
}

func generateCapitalQuizzes(variants, amount int) []CapitalQuiz {
	quizes := make([]CapitalQuiz, 0, amount)

	for i := 0; i < amount; i += 0 {
		quiz := generateCapitalQuiz(variants)
		isNew := true

		for _, q := range quizes {
			if q.Answer == quiz.Answer {
				isNew = false
				break
			}
		}

		if isNew {
			quizes = append(quizes, quiz)
			i++
		}
	}

	return quizes
}

func generateCountryQuizzes(variants, amount int) []CountryQuiz {
	quizzes := make([]CountryQuiz, 0, amount)

	for i := 0; i < amount; i += 0 {
		quiz := generateCountryQuiz(variants)
		isNew := true

		for _, q := range quizzes {
			if q.Answer == quiz.Answer {
				isNew = false
				break
			}
		}

		if isNew {
			quizzes = append(quizzes, quiz)
			i++
		}
	}

	return quizzes
}
