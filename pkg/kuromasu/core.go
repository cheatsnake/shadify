package kuromasu

import "errors"

func Generate(w, h, fill int) (Core, error) {
	if w > maxSideLength || h > maxSideLength {
		return Core{}, errors.New(failedMaxSideLength)
	}

	if w < minSideLength || h < minSideLength {
		return Core{}, errors.New(failedMinSideLength)
	}

	if fill > maxFillPercent || fill < minFillPercent {
		return Core{}, errors.New(failedFillPercent)
	}

	result := generator(w, h, fill)
	return result, nil
}

func Verify(sol [][]string) (VerifyResult, error) {
	if len(sol) > maxSideLength || len(sol[0]) > maxSideLength {
		return VerifyResult{}, errors.New(failedMaxSideLength)
	}

	if len(sol) < minSideLength || len(sol[0]) < minSideLength {
		return VerifyResult{}, errors.New(failedMinSideLength)
	}

	result := verifier(sol)
	return result, nil
}
