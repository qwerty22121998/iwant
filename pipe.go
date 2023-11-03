package iwant

func Map[F any, T any](data []F, fn func(f F) T) []T {
	result := make([]T, len(data))
	for i, d := range data {
		result[i] = fn(d)
	}
	return result
}

func Reduce[E any, R any](data []E, fn func(r R, e E) R, init R) R {
	result := init
	for _, d := range data {
		result = fn(result, d)
	}
	return result
}

func Filter[T any](data []T, fn func(T) bool) []T {
	var result []T
	for _, v := range data {
		if fn(v) {
			result = append(result, v)
		}
	}
	return result
}
