package day12

import "testing"

func TestCountMaps2(t *testing.T) {
	//answers := []float64{6, 4, 0, 6}
	answers := []float64{0}
	//result1, _ := countNumbers([]byte("[1,2,3]"), 2)
	//result2, _ := countNumbers([]byte(`[1,{"c":"red","b":2},3]`), 2)
	result3, _ := countNumbers([]byte(`{"d":"red","e":[1,2,3,4],"f":5}`), 2)
	//result4, _ := countNumbers([]byte(`[1,"red",5]`), 2)
	//results := []float64{result1, result2, result3, result4}
	results := []float64{result3}

	for index, answer := range answers {
		if answer != results[index] {
			t.Fatalf("the answer was: %f, and got: %f", answer, results[index])
		}
	}
}
