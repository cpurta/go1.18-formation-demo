package sums

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSumsIntsAndFloats(t *testing.T) {
	ints := map[string]int64{
		"first":  34,
		"second": 12,
	}

	// Initialize a map for the float values
	floats := map[string]float64{
		"first":  35.98,
		"second": 26.99,
	}

	require.Equal(t, int64(46), SumInts(ints))
	require.Equal(t, float64(62.97), SumFloats(floats))
	require.Equal(t, float64(108.97), SumNumbers(ints, floats))

	// Test using the Number interface map

	numbers := map[string]NonGenericNumber{
		"first":  Int(34),
		"second": Float(35.98),
		"third":  Float(26.99),
		"fourth": Int(12),
	}

	require.Equal(t, "108.97", fmt.Sprintf("%.2f", SumInterfaceNumbers(numbers)))
}
