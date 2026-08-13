package day2

import (
	"testing"
)

func TestCalcSurfaceArea(t *testing.T) {
	wantend_results := [2]uint32{58, 43}
	got_results := [2]uint32{CalcSurfaceArea(2, 3, 4), CalcSurfaceArea(1, 1, 10)}

	for index, result := range got_results {
		got := wantend_results[index]
		if result != got {
			t.Fatalf("got: %d, want: %d", got, result)
		}
	}
}

func TestRibonLength(t *testing.T) {
	wantend_results := [2]uint32{34, 14}
	got_results := [2]uint32{CalcRibonLength(2, 3, 4), CalcRibonLength(1, 1, 10)}

	for index, got := range got_results {
		result := wantend_results[index]
		if result != got {
			t.Fatalf("got: %d, want: %d", got, result)
		}
	}
}
