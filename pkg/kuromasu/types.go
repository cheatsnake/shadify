package kuromasu

type Core struct {
	Width  int        `json:"width"`
	Height int        `json:"height"`
	Task   [][]string `json:"task"`
	Grid   [][]string `json:"grid"`
}

type position struct {
	Row    int `json:"row"`
	Column int `json:"column"`
}

type VerifyResult struct {
	IsValid  bool     `json:"isValid"`
	Position position `json:"position"`
}
