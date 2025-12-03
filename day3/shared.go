package day3

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type bank []int

func readBanks(filename string) []bank {
	var banks []bank

	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		var currentBank bank

		for _, number := range strings.Split(line, "") {
			joltage, err := strconv.Atoi(number)

			if err != nil {
				panic(err)
			}

			currentBank = append(currentBank, joltage)
		}

		banks = append(banks, currentBank)
	}

	return banks
}
