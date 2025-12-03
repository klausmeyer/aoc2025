package day3

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestEvaluateBankMethodTwo(t *testing.T) {
	var require = require.New(t)

	require.Equal(987654321111, evaluateBankMethodTwo(bank{9, 8, 7, 6, 5, 4, 3, 2, 1, 1, 1, 1, 1, 1, 1}))
	require.Equal(811111111119, evaluateBankMethodTwo(bank{8, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 9}))
	require.Equal(434234234278, evaluateBankMethodTwo(bank{2, 3, 4, 2, 3, 4, 2, 3, 4, 2, 3, 4, 2, 7, 8}))
}

func TestEvaluateBanksMethodTwoWithExampleInput(t *testing.T) {
	var banks = readBanks("inputs/example.txt")

	var result = evaluateBanksMethodTwo(banks)

	require.Equal(t, 3121910778619, result)
}

func TestEvaluateBanksMethodTwoWithRealInput(t *testing.T) {
	var banks = readBanks("inputs/input.txt")

	var result = evaluateBanksMethodTwo(banks)

	require.Equal(t, 173300819005913, result)
}
