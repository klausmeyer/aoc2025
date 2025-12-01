package day1

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestReadSequence(t *testing.T) {
	require := require.New(t)

	var sequence = ReadSequence("inputs/example.txt")

	require.Len(sequence, 10)

	require.Equal(DirectionLeft, sequence[0].direction)
	require.Equal(68, sequence[0].steps)

	require.Equal(DirectionLeft, sequence[1].direction)
	require.Equal(30, sequence[1].steps)

	require.Equal(DirectionRight, sequence[2].direction)
	require.Equal(48, sequence[2].steps)
}
