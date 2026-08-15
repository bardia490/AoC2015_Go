package day4

import "testing"

func TestFindingTheRightIntegerForHash(t *testing.T) {
	want_results := [2]int{609043, 1048970}
	got_results := [2]int{FindCorrectHash1([]byte("abcdef")), FindCorrectHash1([]byte("pqrstuv"))}

	for index, got := range got_results {
		result := want_results[index]
		if result != got {
			t.Fatalf("got: %d, want: %d", got, result)
		}
	}
}
