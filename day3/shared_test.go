package day3

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestReadBanks(t *testing.T) {
	var require = require.New(t)

	var banks = readBanks("inputs/example.txt")

	require.Len(banks, 4)
	require.Equal(bank{9, 8, 7, 6, 5, 4, 3, 2, 1, 1, 1, 1, 1, 1, 1}, banks[0])
}
