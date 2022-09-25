package camp

type Core struct {
	Width       int     `json:"width"`
	Height      int     `json:"height"`
	Trees       int     `json:"trees"`
	RowTents    []int   `json:"rowTents"`
	ColumnTents []int   `json:"columnTents"`
	Task       [][]int `json:"task"`
	Solution [][]int `json:"solution"`
}

type VerifyResult struct {
	IsError  bool   `json:"isError"`
	Position string `json:"position"`
}
