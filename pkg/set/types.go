package set

type Core struct {
	FreeCards    []Card   `json:"freeCards"`
	Layout       []Card   `json:"layout"`
	PossibleSets [][]Card `json:"possibleSets"`
	WonCards     []Card   `json:"wonCards"`
	State        string   `json:"state"`
}

type Card struct {
	Id      int    `json:"_id"`
	Number  int    `json:"number"`
	Shape   string `json:"shape"`
	Color   string `json:"color"`
	Shading string `json:"shading"`
}
