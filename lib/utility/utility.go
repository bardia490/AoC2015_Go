package utility

import (
	"bufio"
	"bytes"
	"io"
	"iter"
)

// Source - https://stackoverflow.com/a/52153000
// Posted by Daniel Castillo, modified by community. See post 'Timeline' for change history
// Retrieved 2026-08-17, License - CC BY-SA 4.0

func LineCounter(r io.Reader) (int, error) {
	var count int
	const lineBreak = '\n'

	buf := make([]byte, bufio.MaxScanTokenSize)

	for {
		bufferSize, err := r.Read(buf)
		if err != nil && err != io.EOF {
			return 0, err
		}

		var buffPosition int
		for {
			i := bytes.IndexByte(buf[buffPosition:], lineBreak)
			if i == -1 || bufferSize == buffPosition {
				break
			}
			buffPosition += i + 1
			count++
		}
		if err == io.EOF {
			break
		}
	}

	return count, nil
}

// Heap's algorithm
func Generate[T any](input []T) iter.Seq[[]T] {
	return func(yield func([]T) bool) {
		// Work on a copy so generating permutations does not modify input.
		items := append([]T(nil), input...)

		n := len(items)
		counters := make([]int, n)

		// Emit a copy because items is mutated for subsequent permutations.
		emit := func() bool {
			permutation := append([]T(nil), items...)
			return yield(permutation)
		}

		// The empty input has one permutation: [].
		if !emit() {
			return
		}

		i := 0
		for i < n {
			if counters[i] < i {
				if i%2 == 0 {
					items[0], items[i] = items[i], items[0]
				} else {
					items[counters[i]], items[i] =
						items[i], items[counters[i]]
				}

				if !emit() {
					return // Consumer requested termination.

				}

				counters[i]++
				i = 0
			} else {
				counters[i] = 0
				i++
			}
		}
	}
}

// used for numbers that are not bigger than MAX_INT
// if the number is bigger than MAX_INT then it will panic
func Factorial(num uint64) uint64 {
	var result uint64 = 0
	if num == 1 || num == 2 {
		return num
	}

	var counter uint64 = 2

	for ; counter <= num; counter++ {
		result *= counter
	}
	return result
}

// generic Optional value
type Option[T any] struct {
	value T
	ok    bool
}

type Number interface {
	int | int64 | int32 | float64 | float32
}

func SumSlice[T Number](values []T) T {
	var result T = 0
	for _, value := range values {
		result += value
	}
	return result
}

func GenerateSubsets[T any](set []T) iter.Seq[[]T] {
	return generateSubsetsAux([]T{}, set)
}

func generateSubsetsAux[T any](current []T, rest []T) iter.Seq[[]T] {
	return func(yield func([]T) bool) {
		for index, elem := range rest {
			new := append(current, elem)
			if !yield(new) {
				return
			}
			generateSubsetsAux(new, rest[index+1:])(yield)
		}
	}
}
