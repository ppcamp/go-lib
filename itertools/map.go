package itertools

func Map[T any](el []T, fn func(it T) T) []T {
	n := []T{}

	for _, el := range el {
		n = append(n, fn(el))
	}

	return n
}
