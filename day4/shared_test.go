package day4

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestReadMap(t *testing.T) {
	var require = require.New(t)

	var rollMap = readMap("inputs/example.txt")

	require.Len(rollMap, 10)
	require.Equal([]rune{'.', '.', '@', '@', '.', '@', '@', '@', '@', '.'}, rollMap[0])
}
