package day13

import (
	"Aoc2015/lib/set"
	"Aoc2015/lib/utility"
	"math"
	"testing"
)

func TestPart1(t *testing.T) {
	input := [12]string{
		"Alice would gain 54 happiness units by sitting next to Bob.",
		"Alice would lose 79 happiness units by sitting next to Carol.",
		"Alice would lose 2 happiness units by sitting next to David.",
		"Bob would gain 83 happiness units by sitting next to Alice.",
		"Bob would lose 7 happiness units by sitting next to Carol.",
		"Bob would lose 63 happiness units by sitting next to David.",
		"Carol would lose 62 happiness units by sitting next to Alice.",
		"Carol would gain 60 happiness units by sitting next to Bob.",
		"Carol would gain 55 happiness units by sitting next to David.",
		"David would gain 46 happiness units by sitting next to Alice.",
		"David would lose 7 happiness units by sitting next to Bob.",
		"David would gain 41 happiness units by sitting next to Carol.",
	}

	happiness_diagram := make(map[[2]string]int, 12)
	people := set.Create[string](4)

	for _, line := range input {
		generateHappinessDiagram(happiness_diagram, people, line)
	}

	maximum := math.MinInt
	for permutation := range utility.Generate(people.ToSlice()) { // generate the permutation of peoples names
		result := 0
		for index := 0; index < len(permutation)-1; index += 1 {
			result += happiness_diagram[[2]string{permutation[index], permutation[index+1]}]
			result += happiness_diagram[[2]string{permutation[index+1], permutation[index]}]
		}
		result += happiness_diagram[[2]string{permutation[0], permutation[len(permutation)-1]}]
		result += happiness_diagram[[2]string{permutation[len(permutation)-1], permutation[0]}]
		if result > maximum {
			maximum = result
		}
	}

	if maximum != 330 {
		t.Fatalf("the answer was: %d, and got: %d", 330, maximum)
	}
}
