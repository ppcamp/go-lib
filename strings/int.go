package strings

import "strconv"

func ToInt[T int | int32 | int64](v string) T {
	a, err := strconv.Atoi(v)
	if err != nil {
		panic(err)
	}
	return T(a)
}
