package day4

import (
	"bufio"
	"os"
)

func readMap(filename string) [][]rune {
	var rollMap [][]rune

	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		var row []rune

		line := scanner.Text()

		for _, spot := range line {
			row = append(row, spot)
		}

		rollMap = append(rollMap, row)
	}

	return rollMap
}
