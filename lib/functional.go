package lib

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
