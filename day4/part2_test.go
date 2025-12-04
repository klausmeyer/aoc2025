package day4

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestEvaluateMapPartTwoWithExampleInput(t *testing.T) {
	var rollMap = readMap("inputs/example.txt")

	var result = evaluateMapPartTwo(rollMap)

	require.Equal(t, 43, result)
}

func TestEvaluateMapPartTwoWithRealInput(t *testing.T) {
	var rollMap = readMap("inputs/input.txt")

	var result = evaluateMapPartTwo(rollMap)

	require.Equal(t, 8910, result)
}
