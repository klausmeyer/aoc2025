package day4

func removeRolls(rollMap [][]rune) int {
	var rollsRemoved int

	for row := 0; row < len(rollMap); row++ {
		for col := 0; col < len(rollMap[row]); col++ {
			if rollMap[row][col] != roll {
				continue
			}

			if countAdjacentPositions(rollMap, row, col) <= maxRolls {
				rollMap[row][col] = 'x'
				rollsRemoved++
			}
		}
	}

	if rollsRemoved > 0 {
		rollsRemoved += removeRolls(rollMap)
	}

	return rollsRemoved
}

func evaluateMapPartTwo(rollMap [][]rune) int {
	var rollsRemoved = removeRolls(rollMap)

	return rollsRemoved
}
