package day3

func evaluateBanks(banks []bank) int {
	var total int

	for _, bank := range banks {
		var sum int

		for i, left := range bank {
			for j := i + 1; j < len(bank); j++ {
				var right = bank[j]

				var tempSum = (left * 10) + right

				if tempSum > sum {
					sum = tempSum
				}
			}
		}

		total += sum
	}

	return total
}
