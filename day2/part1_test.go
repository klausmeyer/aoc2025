package day2

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestEvaluateRangesWithExampleInput(t *testing.T) {
	var ranges = ReadRanges("inputs/example.txt")

	var result = EvaluateRanges(ranges)

	require.Equal(t, 1227775554, result)
}

func TestEvaluateRangesWithRealInput(t *testing.T) {
	var ranges = ReadRanges("inputs/input.txt")

	var result = EvaluateRanges(ranges)

	require.Equal(t, 28146997880, result)
}
