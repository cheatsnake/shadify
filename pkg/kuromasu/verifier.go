package kuromasu

func verifier(sol [][]string) VerifyResult {
	test := make([][]string, len(sol))

	for i := range test {
		row := make([]string, len(sol[0]))
		copy(row, sol[i])
		test[i] = row
	}

	for i := range test {
		for j := range test[i] {
			if test[i][j] != boxCell {
				test[i][j] = calcViewCount(j, i, &test)
			}

			if sol[i][j] != emptyCell {
				isEqual := sol[i][j] == test[i][j]

				if !isEqual {
					return VerifyResult{
						IsValid:  false,
						Position: position{Row: i + 1, Column: j + 1}}
				}
			}
		}
	}

	return VerifyResult{IsValid: true}
}
