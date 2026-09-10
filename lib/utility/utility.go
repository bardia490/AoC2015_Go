package utility

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"iter"
	"math"
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
// this algorithm works creates a copy of the original input
// hence it does not modify the original input slice and the slice
// can be used everywhere else
func GeneratePermutations[T any](input []T) iter.Seq[[]T] {
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

// generates all the subsets of an slice in a non deterministic order
func GenerateSubsets[T any](set []T) iter.Seq[[]T] {
	return auxGenerateSubsets([]T{}, set)
}

func auxGenerateSubsets[T any](current []T, rest []T) iter.Seq[[]T] {
	return func(yield func([]T) bool) {
		for index, elem := range rest {
			new := append(current, elem)
			if !yield(new) {
				return
			}
			auxGenerateSubsets(new, rest[index+1:])(yield)
		}
	}
}

// advance index
func auxGenerateSubsetsNAdvanceIndex(current_indexes []int, max_index int) bool {
	last_index := len(current_indexes) - 2 // take the element one from the last
	movement := 0
	for last_index > -1 {
		if current_indexes[last_index] < max_index-movement-1 {
			current_indexes[last_index] += 1
			break
		}
		last_index--
		movement += 1
	}
	// if we the last_index was less than zero means we have reached the end of the array
	if last_index < 0 {
		return false
	}
	for ; last_index < len(current_indexes)-1; last_index++ {
		current_indexes[last_index+1] = current_indexes[last_index] + 1
	}
	return true
}

// copies the elements from array [source] to [dest] with specified indexes from [indexes]
// panics if any index is out of bounds in either array
func CopyWithIndexes[T any](dest, source []T, indexes []int) {
	for dest_index, source_index := range indexes {
		dest[dest_index] = source[source_index]
	}
}

// count is the number of elements in each subset
// error will probably be removed in the future
func GenerateSubsetsN[T any](set []T, count int) (iter.Seq[[]T], error) {
	if count < 2 {
		return nil, fmt.Errorf("length of the subset cannot be less than 2")
	}
	// set up the place-holder array
	subset := make([]T, count)
	// set up an array for keep track of the current indexes, initially it holds the indexes from 0 to N-2
	current_indexes := make([]int, count)
	// fill the place holder with the first batch of [set[0], set[1], ..., set[N-2]]
	for index := range count {
		current_indexes[index] = index
	}
	// for the first iteration of the for loop we need to decrement the last index
	current_indexes[count-2] -= 1
	// to keep track of the last spot in the subset array
	last_index := count - 1
	number_of_elements := len(set)

	return func(yield func([]T) bool) {
		for auxGenerateSubsetsNAdvanceIndex(current_indexes, number_of_elements-1) {
			CopyWithIndexes(subset, set, current_indexes)
			for loop_index := current_indexes[last_index]; loop_index < number_of_elements; loop_index++ {
				subset[last_index] = set[loop_index]
				if !yield(subset) {
					return
				}
			}
		}
	}, nil
}

// e.g. GenerateNumberDivisors(8) => 1, 2, 4, 8
func GenerateNumberDivisors(num int) iter.Seq[int] {
	return func(yield func(int) bool) {
		if num <= 0 {
			return
		}

		num_sqrt := int(math.Sqrt(float64(num))) + 1
		for n := 1; n < num_sqrt; n++ {
			if num%n == 0 {
				if !yield(n) {
					return
				}
				other_divisor := num / n
				if n != other_divisor {
					if !yield(other_divisor) {
						return
					}
				}
			}
		}

		//if !yield(num) {
		//	return
		//}
	}
}

// this is slower than the GenerateNumberDivisors but will return the numbers sequentually
func GenerateNumberDivisorsSorted(num int) iter.Seq[int] {
	return func(yield func(int) bool) {
		if num <= 0 {
			return
		}

		half_num := num / 2
		for n := 1; n < half_num; n++ {
			if num%n == 0 {
				if !yield(n) {
					return
				}
			}
		}

		if !yield(num) {
			return
		}
	}
}

func Abs[T Number](num1 T) T {
	if num1 >= 0 {
		return num1
	}
	return -num1
}

// converts bytes to integerts
// returns error in case of empty slice
// returns error in case one of the components isn't a correct ascii character
// numbers can have one minus or one plus behind them (e.g +23, -5)
func BytesToIntSigned(s []byte) (int, error) {
	if len(s) == 0 {
		return 0, fmt.Errorf("length of the slice was empty, conversion faild")
	}

	sign := 1
	i := 0
	switch s[0] {
	case '-':
		sign = -1
		i = 1
		if len(s) == 1 {
			return 0, fmt.Errorf("length of the slice was empty, conversion faild")
		}
	case '+':
		i = 1
		if len(s) == 1 {
			return 0, fmt.Errorf("length of the slice was empty, conversion faild")
		}
	}

	n := 0
	for ; i < len(s); i++ {
		c := s[i]
		if c < '0' || c > '9' {
			return 0, fmt.Errorf("could not convert character to integer: %c", c)
		}
		n = n*10 + int(c-'0')
	}
	return sign * n, nil
}
