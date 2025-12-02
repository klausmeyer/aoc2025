package day2

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIsInvalidProductMethodTwo(t *testing.T) {
	var require = require.New(t)

	require.True(isInvalidProductMethodTwo(123123))
}

func TestEvaluateRangesMethodTwoWithExampleInput(t *testing.T) {
	var ranges = ReadRanges("inputs/example.txt")

	var result = EvaluateRangesMethodTwo(ranges)

	require.Equal(t, 4174379265, result)
}

func TestEvaluateRangesMethodTwoWithRealInput(t *testing.T) {
	var ranges = ReadRanges("inputs/input.txt")

	var result = EvaluateRangesMethodTwo(ranges)

	require.Equal(t, 40028128307, result)
}
