package utility

import "testing"

func TestAuxGenerateSubsetsNAdvanceIndex(t *testing.T) {
	s := []int{0, 1, 2, 5}
	result := auxGenerateSubsetsNAdvanceIndex(s, 4)
	if !result {
		t.Fatal("result is false in test1")
	}

	answer := []int{0, 1, 3, 4}
	for index, element := range s {
		if element != answer[index] {
			t.Fatalf("at index %d, element was: %d, but answer was: %d", index, element, answer[index])
		}
	}
	s = []int{0, 1, 4, 5}
	result = auxGenerateSubsetsNAdvanceIndex(s, 4)
	if !result {
		t.Fatal("result is false in test2")
	}

	answer = []int{0, 2, 3, 4}
	for index, element := range s {
		if element != answer[index] {
			t.Fatalf("at index %d, element was: %d, but answer was: %d", index, element, answer[index])
		}
	}

	s = []int{2, 3, 4, 5}
	result = auxGenerateSubsetsNAdvanceIndex(s, 4)
	if result {
		t.Fatal("result is true in test3")
	}

	s = []int{0, 1, 1, 1}
	result = auxGenerateSubsetsNAdvanceIndex(s, 5)
	if !result {
		t.Fatal("result is false in test2")
	}

	answer = []int{0, 1, 2, 3}
	for index, element := range s {
		if element != answer[index] {
			t.Fatalf("at index %d, element was: %d, but answer was: %d", index, element, answer[index])
		}
	}
}

func TestCopyWithIndexes(t *testing.T) {
}
