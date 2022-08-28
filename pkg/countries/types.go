package countries

// Country data from the database
type CountryData struct {
	Code    string `json:"code"`    // ISO 3166-1 alpha-2 code of the country
	Country string `json:"country"` // The name of the country
	Capital string `json:"capital"` // The name of the capital of the country
}

// Data for the country quiz
type CountryQuiz struct {
	Flag     string   `json:"flag"`     // Link to the image of the country's flag
	Variants []string `json:"variants"` // Answer options
	Answer   string   `json:"answer"`   // Correct answer
}

// Data for the capital quiz
type CapitalQuiz struct {
	Country  string   `json:"country"`  // The name of the country
	Flag     string   `json:"flag"`     // Link to the image of the country's flag
	Variants []string `json:"variants"` // Answer options
	Answer   string   `json:"answer"`   // Correct answer
}
