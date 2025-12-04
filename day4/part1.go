package day4

const roll = '@'
const maxRolls = 3 // fewer than 4

func safeCheckPosition(rollMap [][]rune, row, col int) bool {
	if row < 0 || row >= len(rollMap) {
		return false
	}

	if col < 0 || col >= len(rollMap[row]) {
		return false
	}

	return rollMap[row][col] == roll
}

func countAdjacentPositions(rollMap [][]rune, row, col int) int {
	var count int

	// 1 2 3
	// 4   5
	// 6 7 8

	// 1 - Top left
	if safeCheckPosition(rollMap, row-1, col-1) {
		count++
	}

	// 2 - Top Middle
	if safeCheckPosition(rollMap, row-1, col) {
		count++
	}

	// 3 - Top Right
	if safeCheckPosition(rollMap, row-1, col+1) {
		count++
	}

	// 4 - Left
	if safeCheckPosition(rollMap, row, col-1) {
		count++
	}

	// 5 - Right
	if safeCheckPosition(rollMap, row, col+1) {
		count++
	}

	// 6 - Bottom Left
	if safeCheckPosition(rollMap, row+1, col-1) {
		count++
	}

	// 7 - Bottom Middle
	if safeCheckPosition(rollMap, row+1, col) {
		count++
	}
	// 8 - Bottom Right
	if safeCheckPosition(rollMap, row+1, col+1) {
		count++
	}

	return count
}

func evaluateMap(rollMap [][]rune) int {
	var numberRolls int

	for row := 0; row < len(rollMap); row++ {
		for col := 0; col < len(rollMap[row]); col++ {
			if rollMap[row][col] != roll {
				continue
			}

			if countAdjacentPositions(rollMap, row, col) <= maxRolls {
				numberRolls++
			}
		}
	}

	return numberRolls
}
