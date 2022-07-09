package strings

import "strconv"

func ToFloat[T float32 | float64](v string) (T, error) {
	if v == "" {
		v = "0"
	}

	a, err := strconv.ParseFloat(v, 64)
	if err != nil {
		var r T
		return r, err
	}
	return T(a), nil
}
