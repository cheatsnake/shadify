package camp

func Generate(w, h, t int) Core {
	field := generateField(w, h, t)
	rowTents, columnTents := countTents(field)
	task := hideTents(field)

	return Core{w, h, t, rowTents, columnTents, task, field}
}

func Verify(field [][]int) (VerifyResult, error) {
	return VerifyResult{}, nil
}
