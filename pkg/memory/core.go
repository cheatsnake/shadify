package memory

func Generate(w, h, ps int, showPositions bool) (Core, error) {
	grid, positions, err := generator(w, h, ps, showPositions)
	if err != nil {
		return Core{}, err
	}

	return Core{
		Width:         w,
		Height:        h,
		PairSize:      ps,
		TotalPairs:    (w * h) / ps,
		Grid:          grid,
		PairPositions: positions,
	}, nil
}
