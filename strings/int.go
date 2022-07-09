package strings

import "strconv"

func ToInt[T int | int32 | int64](v string) (T, error) {
	if v == "" {
		v = "0"
	}

	a, err := strconv.Atoi(v)
	if err != nil {
		var r T
		return r, err
	}
	return T(a), nil
}
