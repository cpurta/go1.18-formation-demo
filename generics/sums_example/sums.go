package sums

func SumInts(m map[string]int64) int64 {
	var s int64
	for _, v := range m {
		s += v
	}
	return s
}

func SumFloats(m map[string]float64) float64 {
	var s float64
	for _, v := range m {
		s += v
	}
	return s
}

func SumNumbers(ints map[string]int64, floats map[string]float64) float64 {
	return float64(SumInts(ints)) + SumFloats(floats)
}

// Using an interface to get around not having a generics

type NonGenericNumber interface {
	Number() float64
}

type Int int
type Float float64

func (i Int) Number() float64 {
	return float64(i)
}

func (f Float) Number() float64 {
	return float64(f)
}

func SumInterfaceNumbers(m map[string]NonGenericNumber) float64 {
	var s float64
	for _, v := range m {
		s += v.Number()
	}
	return s
}
