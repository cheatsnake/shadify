package camp

import (
	"github.com/cheatsnake/shadify/internal/helpers"
)

func GenerateField(w, h, t int) [][]int {
	field := make([][]int, h)
	for i := range field {
		field[i] = make([]int, w)
	}

	for t > 0 {
		treeY := helpers.GetRandomInteger(0, h-1)
		treeX := helpers.GetRandomInteger(0, w-1)

		campX, campY, isCampFit := placeCamp(treeX, treeX, field)
		if (!isCampFit) {
			continue
		}

		if (field[treeY][treeX] == 0) {
			field[treeY][treeX] = treeValue
			field[campY][campX] = campValue
			t--
		}
	}

	return field
}

func placeCamp(x, y int, field [][]int) (int, int, bool) {
	possibleCoords := [][]int{{x-1, y}, {x+1, y}, {x, y-1}, {x, y+1}}
	randCoords := possibleCoords[helpers.GetRandomInteger(0, len(possibleCoords) - 1)]
	campX, campY := randCoords[0], randCoords[1]


	isCoordsFit := checkIsCoordsFit(campX, campY, field)
	isCampFit := checkIsCampFit(campX, campY, field)

	// println(isCampFit, isCampFit)

	return campX, campY, isCoordsFit && isCampFit
}

func checkIsCoordsFit(x, y int, field [][]int) bool {
	return x >= 0 && y >= 0 && x < len(field[0]) && y < len(field)
}

func checkIsCampFit(campX, campY int, field [][]int) (bool) {
	treesCells := 0
	campCells := 0

	checkCampArea := func (x, y int) {
		if (checkIsCoordsFit(x, y, field)) {
			if field[y][x] == treeValue {
				treesCells++
			}
			if field[y][x] == campValue {
				campCells++
			}
		}
	}

	checkCampArea(campX-1, campY)
	checkCampArea(campX+1, campY)
	checkCampArea(campX, campY-1)
	checkCampArea(campX, campY+1)

	return treesCells == 1 && campCells == 0
}