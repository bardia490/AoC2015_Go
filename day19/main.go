package day19

import (
	"Aoc2015/lib/set"
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

type Drug struct {
	initial     string
	replacement string
	single_atom bool
}

func generateDrugsArray(drugs []Drug, in string, index int) {
	contents := strings.Split(in, " ") // from the input we know that this will only have 3 elements
	single_atom := false
	if len(contents[0]) == 1 {
		single_atom = true
	}

	drugs[index] = Drug{
		initial:     contents[0],
		replacement: contents[2],
		single_atom: single_atom,
	}
}

func part1(drugs []Drug, initial_drug []byte) int {
	drugs_set := set.Create[string](50)

	for _, drug := range drugs {
		if drug.single_atom {
			for index, atom := range initial_drug {
				start_index := 0
				end_index := 0
				if drug.initial[0] == atom {
					start_index = index
					end_index = index + 1
				}
				if start_index != end_index {
					la := len(initial_drug[0:start_index])
					lreplacement := len(drug.replacement)
					c := make([]byte, la, la+len(initial_drug[end_index:])+lreplacement)
					_ = copy(c, initial_drug[0:start_index])
					c = append(c, drug.replacement...)
					c = append(c, initial_drug[end_index:]...)
					drugs_set.Append(string(c))
				}
			}
		} else {
			for index := 0; index < len(initial_drug)-1; index += 1 {
				start_index := 0
				end_index := 0
				if drug.initial == string(initial_drug[index:index+2]) {
					start_index = index
					end_index = index + 2
				}
				if start_index != end_index {
					la := len(initial_drug[0:start_index])
					lreplacement := len(drug.replacement)
					c := make([]byte, la, la+len(initial_drug[end_index:])+lreplacement)
					_ = copy(c, initial_drug[0:start_index])
					c = append(c, drug.replacement...)
					c = append(c, initial_drug[end_index:]...)
					drugs_set.Append(string(c))
				}
			}
		}
	}
	return len(drugs_set)
}

func part2(drugs []Drug, initial_drug []byte) int {
	result := 0
	return result
}

func Solution1(f *os.File) {
	sc := bufio.NewScanner(f)

	index := 0
	var drugs [43]Drug
	for sc.Scan() {
		line := sc.Text()
		if len(line) == 0 {
			break
		}
		generateDrugsArray(drugs[:], line, index)
		index += 1
	}
	sc.Scan()
	initial_drug := sc.Bytes()
	result := part1(drugs[:], initial_drug)

	if err := sc.Err(); err != nil {
		panic(fmt.Sprintf("there was a problem reading the file: %s", err.Error()))
	}

	fmt.Println("the solution to day17 part 1 is:", result)
	f.Seek(0, io.SeekStart)
}

func Solution2(f *os.File) {
	sc := bufio.NewScanner(f)

	index := 0
	var drugs [43]Drug
	for sc.Scan() {
		line := sc.Text()
		if len(line) == 0 {
			break
		}
		generateDrugsArray(drugs[:], line, index)
		index += 1
	}
	sc.Scan()
	initial_drug := sc.Bytes()
	result := part2(drugs[:], initial_drug)

	if err := sc.Err(); err != nil {
		panic(fmt.Sprintf("there was a problem reading the file: %s", err.Error()))
	}

	fmt.Println("the solution to day17 part 2 is:", result)
}
