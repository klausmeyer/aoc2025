package day2

import (
	"strconv"
)

func isInvalidProduct(number int) bool {
	var numberAsString = strconv.Itoa(number)
	var length = len(numberAsString)

	if (length % 2) != 0 {
		return false
	}

	firstHalf := numberAsString[:length/2]
	secondHalf := numberAsString[length/2:]

	if firstHalf == secondHalf {
		return true
	}

	return false
}

func EvaluateRanges(productRanges []productRange) int {
	var invalid []int

	for _, products := range productRanges {
		for num := products.first; num <= products.last; num++ {
			if isInvalidProduct(num) {
				invalid = append(invalid, num)
			}
		}
	}

	return arraySum(invalid)
}
