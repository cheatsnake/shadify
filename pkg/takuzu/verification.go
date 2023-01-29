package takuzu

import (
	"errors"
	"strconv"
	"strings"
)

func verify(field [][]string) (*VerificationResult, error) {

	if len(field) == 0 {
		return &VerificationResult{}, errors.New(emptyTask)
	}

	rows := []string{}
	columns := []string{}

	// Transform rows & columns into strings
	for i := range field {
		row := strings.Join(field[i][:], "")
		column := ""
		for j := range field {
			column += field[j][i]
		}
		rows = append(rows, row)
		columns = append(columns, column)
	}

	result := checkLines(rows, "row")

	if !result.IsValid {
		return result, errors.New(result.Message)
	}

	result = checkLines(columns, "column")

	if !result.IsValid {
		return result, errors.New(result.Message)
	}

	return result, nil
}

func checkDuplication(lines []string, lineType string) []string {
	position := []string{}
	for i := range lines {
		for j := i + 1; j < len(lines); j++ {
			if lines[i] == lines[j] {
				position = append(
					position,
					lineType+"-"+strconv.Itoa(i+1),
					lineType+"-"+strconv.Itoa(j+1),
				)
			}
		}
		if len(position) > 0 {
			break
		}
	}
	return position
}

func defineVerificationError(msg string, pos []string) *VerificationResult {
	return &VerificationResult{
		IsValid:  false,
		Message:  msg,
		Position: pos,
	}
}

func checkTripleValues(line string) bool {
	result := strings.Contains(line, strings.Repeat(takuzuZero, 3)) ||
		strings.Contains(line, strings.Repeat(takuzuOne, 3))
	return result
}

func checkBalance(line string) bool {
	zeros := strings.Count(line, takuzuZero)
	ones := strings.Count(line, takuzuOne)
	result := zeros == ones
	return result
}

func checkLines(lines []string, lineType string) *VerificationResult {
	result := &VerificationResult{
		IsValid:  true,
		Message:  "",
		Position: []string{},
	}

	duplicationLines := checkDuplication(lines, lineType)
	if len(duplicationLines) > 0 {
		return defineVerificationError("duplication", duplicationLines)
	}

	for i, line := range lines {
		isTripled := checkTripleValues(line)
		if isTripled {
			return defineVerificationError("triple values", []string{lineType + "-" + strconv.Itoa(i+1)})
		}

		isBalanced := checkBalance(line)
		if !isBalanced {
			return defineVerificationError("failed balance", []string{lineType + "-" + strconv.Itoa(i+1)})
		}
	}
	return result
}
