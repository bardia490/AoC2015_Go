package day1

import (
	"local/Aoc2015/lib/functional"
	"testing"
)

func TestAdd(t *testing.T) {
	want := 3
	got := Add(1, 2)

	if want != got {
		t.Fatalf("Hello() = %d, want %d", got, want)
	}
}

func TestMoveFloors(t *testing.T) {
	wantend_results := [5]int{0, 3, 3, -1, -3}
	inputs := [5]string{"(())", "(()(()(", "))(((((", "))(", ")())())"}
	got_results := functional.Map(MoveFloors, inputs[:])

	for index, result := range got_results {
		got := wantend_results[index]
		if result != got {
			t.Fatalf("Hello() = %d, want %d", got, result)
		}
	}
}
