package day9

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

func generateDistanceMap(in string, distances map[[2]string]int, cities set.Set[string]) {
	contents := strings.Split(in, " ")

	key := [2]string{contents[0], contents[2]}
	val, err := strconv.Atoi(contents[4])
	if err != nil {
		panic(fmt.Sprintf("there was an error spliting the string: %s, with this error: %s", contents[4], err.Error()))
	}

	cities.Append(key[0])
	cities.Append(key[1])
	distances[key] = val
}

func Solution1(f *os.File) {

	distances := make(map[[2]string]int, 100)
	cities := set.Create[string](10)

	sc := bufio.NewScanner(f)

	for sc.Scan() {
		line := sc.Text()
		generateDistanceMap(line, distances, cities)
	}

	if err := sc.Err(); err != nil {
		panic(fmt.Sprintf("there was a problem reading the file: %s", err.Error()))
	}

	minimum := math.MaxInt

	for permutation := range utility.Generate(cities.ToSlice()) {
		result := 0
		for index := 0; index < len(permutation)-1; index += 1 {
			key := [2]string{permutation[index], permutation[index+1]}
			if val, ok := distances[key]; ok {
				result += val
			} else {
				val := distances[[2]string{permutation[index+1], permutation[index]}]
				result += val
			}
		}
		if minimum > result {
			minimum = result
		}
	}
	fmt.Println("the solution to day9 part 1 is:", minimum)
	f.Seek(0, io.SeekStart)
}

func Solution2(f *os.File) {

	distances := make(map[[2]string]int, 100)
	cities := set.Create[string](10)

	sc := bufio.NewScanner(f)

	for sc.Scan() {
		line := sc.Text()
		generateDistanceMap(line, distances, cities)
	}

	if err := sc.Err(); err != nil {
		panic(fmt.Sprintf("there was a problem reading the file: %s", err.Error()))
	}

	maximum := 0

	for permutation := range utility.Generate(cities.ToSlice()) {
		result := 0
		for index := 0; index < len(permutation)-1; index += 1 {
			key := [2]string{permutation[index], permutation[index+1]}
			if val, ok := distances[key]; ok {
				result += val
			} else {
				val := distances[[2]string{permutation[index+1], permutation[index]}]
				result += val
			}
		}
		if maximum < result {
			maximum = result
		}
	}
	fmt.Println("the solution to day9 part 2 is:", maximum)
}
