package kuromasu

func Generate(w, h, fill int) Core {
	result := generator(w, h, fill)
	return result
}

func Verify(sol [][]string) VerifyResult {
	return VerifyResult{}
}
