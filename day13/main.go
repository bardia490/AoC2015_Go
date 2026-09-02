package day13

import (
	"Aoc2015/lib/set"
	"Aoc2015/lib/utility"
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

func generateHappinessDiagram(happiness_diagram map[[2]string]int, people set.Set[string], in string) {
	var map_buffer [2]string

	contents := strings.Split(in, " ")
	map_buffer[0] = contents[0]
	people.Append(contents[0])

	map_buffer[1] = contents[len(contents)-1]
	map_buffer[1] = map_buffer[1][:len(map_buffer[1])-1] // trim the . at the end
	people.Append(map_buffer[1])

	happiness, err := strconv.Atoi(contents[3])
	if err != nil {
		panic(fmt.Sprintf("could not convert the value %s at line: %s to integer, err: %s, terminating program", contents[3], in, err.Error()))
	}

	if contents[2] == "gain" {
		happiness_diagram[map_buffer] = happiness
	} else {
		happiness_diagram[map_buffer] = (-1) * happiness
	}
}

func Solution1(f *os.File) {

	happiness_diagram := make(map[[2]string]int, 56)
	people := set.Create[string](8)

	sc := bufio.NewScanner(f)

	for sc.Scan() {
		line := sc.Text()
		generateHappinessDiagram(happiness_diagram, people, line)
	}

	if err := sc.Err(); err != nil {
		panic(fmt.Sprintf("there was a problem reading the file: %s", err.Error()))
	}

	maximum := math.MinInt
	for permutation := range utility.GeneratePermutations(people.ToSlice()) { // generate the permutation of peoples names
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
	fmt.Println("the solution to day13 part 1 is:", maximum)
	f.Seek(0, io.SeekStart)
}

func Solution2(f *os.File) {
	happiness_diagram := make(map[[2]string]int, 57)
	people := set.Create[string](9)

	sc := bufio.NewScanner(f)

	for sc.Scan() {
		line := sc.Text()
		generateHappinessDiagram(happiness_diagram, people, line)
	}

	people.Append("me") // do not add me to the map since it will return zero anyway

	if err := sc.Err(); err != nil {
		panic(fmt.Sprintf("there was a problem reading the file: %s", err.Error()))
	}

	maximum := math.MinInt

	for permutation := range utility.GeneratePermutations(people.ToSlice()) {
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
	fmt.Println("the solution to day13 part 2 is:", maximum)
}
