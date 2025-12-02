package day2

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type productRange struct {
	first int
	last  int
}

func ReadRanges(filename string) []productRange {
	var ranges []productRange

	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()

	line := scanner.Text()

	items := strings.Split(line, ",")

	for _, item := range items {
		numbers := strings.Split(item, "-")

		first, err := strconv.Atoi(numbers[0])
		if err != nil {
			panic(err)
		}

		last, err := strconv.Atoi(numbers[1])
		if err != nil {
			panic(err)
		}

		ranges = append(ranges, productRange{first: first, last: last})
	}

	return ranges
}

func arraySum(numbers []int) int {
	var sum = 0

	for _, num := range numbers {
		sum += num
	}

	return sum
}
