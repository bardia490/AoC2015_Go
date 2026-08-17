package day5

import "testing"

func TestNaughtyOrNice(t *testing.T) {
	want_results := [5]bool{true, true, false, false, false}
	got_results := [5]bool{NaughtyOrNice("ugknbfddgicrmopn"), NaughtyOrNice("aaa"), NaughtyOrNice("jchzalrnumimnmhp"), NaughtyOrNice("haegwjzuvuyypxyu"), NaughtyOrNice("dvszwmarrgswjxmb")}

	for index, got := range got_results {
		result := want_results[index]
		if result != got {
			t.Fatalf("got: %T, want: %T", got, result)
		}
	}
}

func TestNaughtyOrNice2(t *testing.T) {
	want_results := [4]bool{true, true, false, false}
	got_results := [4]bool{NaughtyOrNice2("qjhvhtzxzqqjkmpb"), NaughtyOrNice2("xxyxx"), NaughtyOrNice2("uurcxstgmygtbstg"), NaughtyOrNice2("ieodomkazucvgmuy")}

	for index, got := range got_results {
		result := want_results[index]
		if result != got {
			t.Fatalf("got: %t, want: %t, at index: %d", got, result, index)
		}
	}
}
