package memory

type Core struct {
	Width         int             `json:"width"`
	Height        int             `json:"height"`
	PairSize      int             `json:"pairSize"`
	TotalPairs    int             `json:"totalPairs"`
	Grid          [][]string      `json:"grid"`
	PairPositions []PairPositions `json:"pairPositions,omitempty"`
}

type PairPositions struct {
	Value     string  `json:"value"`
	Positions [][]int `json:"positions"`
}
