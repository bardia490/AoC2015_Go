package stack

import (
	"errors"
	"fmt"
)

type Stack[T any] []T

var ErrEmptyStack = errors.New("stack is empty")

func Create[T any](count uint32) Stack[T] {
	return make([]T, 0, count)
}

func (s *Stack[T]) Push(elem T) {
	*s = append(*s, elem)
}

func (s *Stack[T]) Pop() (last T, err error) {
	var zero T // the compiler will just zero initialize the value
	if len(*s) == 0 {
		return zero, ErrEmptyStack
	}
	last = (*s)[len(*s)-1]
	*s = (*s)[0 : len(*s)-1]
	return last, nil
}

func (s Stack[T]) Print() {
	for _, v := range s {
		fmt.Println(v)
	}
}

func (s Stack[T]) Len() int {
	return len(s)
}

func (s Stack[T]) Peek() (T, error) {
	var zero T

	if len(s) == 0 {
		return zero, ErrEmptyStack
	}

	return s[len(s)-1], nil
}

func (s Stack[T]) IsEmpty() bool {
	return len(s) == 0
}
