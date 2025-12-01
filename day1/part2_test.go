package day1

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestEvalSequenceMethodTwoWithExampleInput(t *testing.T) {
	var sequence = ReadSequence("inputs/example.txt")

	var result = EvalSequenceMethodTwo(50, sequence)

	require.Equal(t, 6, result)
}

func TestEvalSequenceMethodTwoWithRealInput(t *testing.T) {
	var sequence = ReadSequence("inputs/input.txt")

	var result = EvalSequenceMethodTwo(50, sequence)

	require.Equal(t, 5872, result)
}
