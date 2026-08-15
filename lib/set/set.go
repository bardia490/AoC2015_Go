package set

import "fmt"

type Set[T comparable] map[T]struct{}

// creates an set
// T is the type of the elements inside the set
// count is just the initial size and will we updated accordingly
func Create[T comparable](count uint32) Set[T] { // the reason the return value is struct{}, is because it's zero sized
	return make(map[T]struct{}, count)
}

// this function (for now) will try to append the value if it was present, it doesn't do anything
func (s Set[T]) Append(k T) {
	s[k] = struct{}{}
}

func (s Set[T]) Has(k T) bool {
	_, ok := s[k]
	return ok
}

// this function  will try to remove an item
func (s Set[T]) Remove(k T) {
	delete(s, k)
}

func (s Set[T]) Print() {
	for k := range s {
		fmt.Println(k)
	}
}

func (s Set[T]) Len() int {
	return len(s)
}

// returns a new set that is union of both sets
func Union[T comparable](s1 Set[T], s2 Set[T]) Set[T] {
	len1, len2 := s1.Len(), s2.Len()
	result := Create[T](uint32(max(len1, len2)))

	if s1.Len() <= s2.Len() {
		result = s2
		for k := range s1 {
			if !result.Has(k) {
				result.Append(k)
			}
		}
	} else {
		result = s1
		for k := range s2 {
			if !result.Has(k) {
				result.Append(k)
			}
		}
	}
	return result
}
