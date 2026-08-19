package day10

import "testing"

func TestStepForward(t *testing.T) {
	answers := []string{"11", "21", "1211", "111221", "312211"}
	results := []string{StepForward("1"), StepForward("11"), StepForward("21"), StepForward("1211"), StepForward("111221")}

	for index, answer := range answers {
		if answer != results[index] {
			t.Fatalf("the answer was: %s, and got: %s", answer, results[index])
		}
	}
}

func BenchmarkStepForward(b *testing.B) {
	input := "111221"

	b.ReportAllocs()
	b.ResetTimer()

	for b.Loop() {
		_ = StepForward(input)
	}
}
