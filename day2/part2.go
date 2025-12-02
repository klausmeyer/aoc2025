package day2

import (
	"strconv"
)

func isInvalidProductMethodTwo(number int) bool {
	numberAsString := strconv.Itoa(number)
	length := len(numberAsString)

	for size := 1; size <= length/2; size++ {
		if length%size != 0 {
			continue
		}

		sequence := numberAsString[:size]
		ok := true

		for i := size; i < length; i += size {
			if numberAsString[i:i+size] != sequence {
				ok = false
				break
			}
		}

		if ok {
			return true
		}
	}

	return false
}

func EvaluateRangesMethodTwo(productRanges []productRange) int {
	var invalid []int

	for _, products := range productRanges {
		for num := products.first; num <= products.last; num++ {
			if isInvalidProductMethodTwo(num) {
				invalid = append(invalid, num)
			}
		}
	}

	return arraySum(invalid)
}
