package day2

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestReadRangesWithExampleInput(t *testing.T) {
	var require = require.New(t)

	var ranges = ReadRanges("inputs/example.txt")

	require.Len(ranges, 11)

	require.Equal(ranges[0].first, 11)
	require.Equal(ranges[0].last, 22)

	require.Equal(ranges[1].first, 95)
	require.Equal(ranges[1].last, 115)
}
