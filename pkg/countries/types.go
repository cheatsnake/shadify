package countries

type CountryData struct {
	Code    string `json:"code"`
	Country string `json:"country"`
	Capital string `json:"capital"`
}

type CountryQuiz struct {
	Flag     string   `json:"flag"`
	Variants []string `json:"variants"`
	Answer   string   `json:"answer"`
}

type CapitalQuiz struct {
	Country  string   `json:"country"`
	Flag     string   `json:"flag"`
	Variants []string `json:"variants"`
	Answer   string   `json:"answer"`
}
