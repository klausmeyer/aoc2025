package day3

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestEvaluateBanksWithExampleInput(t *testing.T) {
	var banks = readBanks("inputs/example.txt")

	var result = evaluateBanks(banks)

	require.Equal(t, 357, result)
}

func TestEvaluateBanksWithRealInput(t *testing.T) {
	var banks = readBanks("inputs/input.txt")

	var result = evaluateBanks(banks)

	require.Equal(t, 17452, result)
}
