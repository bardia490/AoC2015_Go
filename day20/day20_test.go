package day20

import (
	"testing"
)

func TestPart1(t *testing.T) {
	answers := [9]int{
		1,
		3,
		4,
		7,
		6,
		12,
		8,
		15,
		13,
	}

	for index, answer := range answers {
		result := calcHousePresents(index + 1)
		if result != answer {
			t.Fatalf("at number: %d the answer was: %d but got %d", index+1, answer, result)
		}
	}

}
