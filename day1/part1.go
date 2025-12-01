package day1

func EvalSequence(start int, sequence []rotation) int {
	var current = start
	var count = 0

	const min = 0
	const max = 99

	for _, rotation := range sequence {

		for i := 0; i < rotation.steps; i++ {
			switch rotation.direction {
			case DirectionLeft:
				current--
			case DirectionRight:
				current++
			}

			if current < min {
				current = max
			}

			if current > max {
				current = min
			}
		}

		if current == 0 {
			count += 1
		}

		// log.Printf("Current: %d\n", current)

	}

	return count
}
