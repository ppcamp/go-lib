package strings

import "strconv"

func ToFloat[T float32 | float64](v string) T {
	a, err := strconv.ParseFloat(v, 64)
	if err != nil {
		panic(err)
	}
	return T(a)
}
