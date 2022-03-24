package sums

// Generic type constraint
type GenericNumber interface {
	int64 | float64 | string
}

func SumGenericNumbers[K comparable, V GenericNumber](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}
