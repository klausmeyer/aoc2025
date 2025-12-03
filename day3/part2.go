package day3

func digitsToInt(digits []int) int {
	var n = 0

	for _, digit := range digits {
		n = n*10 + digit
	}

	return n
}

func evaluateBankMethodTwo(b bank) int {
	const digits = 12

	var remaining = len(b) - digits
	var stack = make([]int, 0, len(b))

	for _, digit := range b {
		for len(stack) > 0 && stack[len(stack)-1] < digit && remaining > 0 {
			stack = stack[:len(stack)-1]
			remaining--
		}

		stack = append(stack, digit)
	}

	return digitsToInt(stack[:digits])
}

func evaluateBanksMethodTwo(banks []bank) int {
	var total int

	for _, bank := range banks {
		total += evaluateBankMethodTwo(bank)
	}

	return total
}
