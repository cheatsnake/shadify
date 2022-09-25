package camp

import (
	"github.com/cheatsnake/shadify/internal/helpers"
)

func generateField(w, h, t int) [][]int {
	i := 0
	field := make([][]int, h)
	for i := range field {
		field[i] = make([]int, w)
	}

	for t > 0 {
		i++
		treeY := helpers.GetRandomInteger(0, h-1)
		treeX := helpers.GetRandomInteger(0, w-1)

		isTreeFit := checkIsAreaFit(treeX, treeY, true, field)
		if !isTreeFit {
			continue
		}

		field[treeY][treeX] = treeValue

		tentX, tentY, isTentFit := placeTent(treeX, treeY, field)
		if !isTentFit {
			field[treeY][treeX] = emptyValue
			continue
		}

		field[tentY][tentX] = tentValue
		t--
	}
	println(i)

	return field
}

func checkIsCoordsFit(x, y int, field [][]int) bool {
	return x >= 0 && y >= 0 && x < len(field[0]) && y < len(field)
}

func placeTent(x, y int, field [][]int) (int, int, bool) {
	possibleCoords := [][]int{{x - 1, y}, {x + 1, y}, {x, y - 1}, {x, y + 1}}
	randCoords := possibleCoords[helpers.GetRandomInteger(0, len(possibleCoords)-1)]
	tentX, tentY := randCoords[0], randCoords[1]

	isCoordsFit := checkIsCoordsFit(tentX, tentY, field)
	if !isCoordsFit {
		return 0, 0, isCoordsFit
	}

	isTentFit := checkIsAreaFit(tentX, tentY, false, field)
	return tentX, tentY, isTentFit
}

func checkIsAreaFit(srcX, srcY int, isTree bool, field [][]int) bool {
	treesAround := 0
	tentsAround := 0
	coordsCount := 0

	if field[srcY][srcX] != emptyValue {
		return false
	}

	checkArea := func(x, y int) {
		if checkIsCoordsFit(x, y, field) {
			coordsCount++
			if field[y][x] == treeValue {
				treesAround++
			}
			if field[y][x] == tentValue {
				tentsAround++
			}
		}
	}

	checkArea(srcX-1, srcY)
	checkArea(srcX+1, srcY)
	checkArea(srcX, srcY-1)
	checkArea(srcX, srcY+1)

	if isTree {
		return treesAround < coordsCount && tentsAround == 0
	}
	return treesAround == 1 && tentsAround == 0
}

func countTents(field [][]int) ([]int, []int) {
	row := make([]int, len(field))
	column := make([]int, len(field[0]))

	for i := range field {
		for j := range field[0] {
			if field[i][j] == tentValue {
				column[j]++
			}
		}
	}

	for i := range field {
		for j := range field[0] {
			if field[i][j] == tentValue {
				row[i]++
			}
		}
	}

	return row, column
}

func hideTents(field [][]int) [][]int {
	task := make([][]int, len(field))
	for i := range task {
		task[i] = make([]int, len(field[0]))
	}

	for i := range field {
		for j := range field[0] {
			if field[i][j] == treeValue {
				task[i][j] = treeValue
			}
		}
	}

	return task
}
