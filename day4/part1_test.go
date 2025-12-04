package day4

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestEvaluateMapWithExampleInput(t *testing.T) {
	var rollMap = readMap("inputs/example.txt")

	var result = evaluateMap(rollMap)

	require.Equal(t, 13, result)
}

func TestEvaluateMapWithRealInput(t *testing.T) {
	var rollMap = readMap("inputs/input.txt")

	var result = evaluateMap(rollMap)

	require.Equal(t, 1474, result)
}
