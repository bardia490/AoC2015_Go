package functional

// create a generic type for the function that turns one input to
// one output
type mappingFunction[T, U any] func(T) U

// create the actual function
// Note: the mappingFunction needs to be followed by [T, U] so the
// compiler can infer the types
func Map[T, U any](f mappingFunction[T, U], in []T) []U {
	result := make([]U, len(in))

	for index, elem := range in {
		result[index] = f(elem)
	}

	return result
}

// the folding function should take a generic variable of type T
// and a generic accumulator of type A
type foldingFunction[A, T any] func(A, T) A

func FoldMap[K comparable, V, A any](f foldingFunction[A, V], acc A, m map[K]V) (accumulator A) {
	for _, elem := range m {
		acc = f(acc, elem)
	}
	return acc
}
