package day1

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestEvalSequenceWithExampleInput(t *testing.T) {
	var sequence = ReadSequence("inputs/example.txt")

	var result = EvalSequence(50, sequence)

	require.Equal(t, 3, result)
}

func TestEvalSequenceWithRealInput(t *testing.T) {
	var sequence = ReadSequence("inputs/input.txt")

	var result = EvalSequence(50, sequence)

	require.Equal(t, 964, result)
}
