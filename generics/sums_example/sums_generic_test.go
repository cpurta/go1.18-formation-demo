package sums

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSumGenericNumbers(t *testing.T) {
	ints := map[string]int64{
		"first":  34,
		"second": 12,
	}

	// Initialize a map for the float values
	floats := map[string]float64{
		"first":  35.98,
		"second": 26.99,
	}

	strings := map[string]string{
		"first":  "34",
		"second": "35.98",
	}

	require.Equal(t, int64(46), SumGenericNumbers(ints))
	require.Equal(t, float64(62.97), SumGenericNumbers(floats))
	require.Equal(t, "35.98", SumGenericNumbers(strings))
	require.Equal(t, float64(108.97), float64(SumGenericNumbers(ints))+SumGenericNumbers(floats))
}
