package memory

func Generate(w, h, ps int, showPositions bool) (Core, error) {
	grid, err := gridGenerator(w, h, ps)
	if err != nil {
		return Core{}, err
	}

	return Core{
		Width:         w,
		Height:        h,
		PairSize:      ps,
		TotalPairs:    (w * h) / ps,
		Grid:          grid,
		PairPositions: []PairPositions{},
	}, nil
}
