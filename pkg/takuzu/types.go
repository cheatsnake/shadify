package takuzu

type Core struct {
	Size  int        `json:"size"`
	Field [][]string `json:"field"`
	Task  [][]string `json:"task"`
}

type VerificationResult struct {
	IsValid  bool     `json:"isValid"`
	Message  string   `json:"message"`
	Position []string `json:"position"`
}
