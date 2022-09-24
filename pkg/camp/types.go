package camp

type Core struct {
	Width       int     `json:"width"`
	Height      int     `json:"height"`
	Trees       int     `json:"trees"`
	RowCamps    []int   `json:"rowCamps"`
	ColumnCamps []int   `json:"columnCamps"`
	Field       [][]int `json:"field"`
}

type VerifyResult struct {
	IsError  bool   `json:"isError"`
	Position string `json:"position"`
}
