package day19

import (
	"Aoc2015/lib/set"
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
)

type Drug struct {
	initial     []byte
	replacement []byte
}

func generateDrugsArray(drugs []Drug, in []byte, index int) {
	contents := bytes.SplitN(in, []byte(" => "), 2)

	drugs[index] = Drug{
		initial:     contents[0],
		replacement: contents[1],
	}
}

func part1(drugs []Drug, initial_drug []byte) int {
	drugs_set := set.Create[string](50)

	for _, drug := range drugs {
		pattern := drug.initial
		pattern_len := len(pattern)

		for i := 0; i < len(initial_drug)-pattern_len; i++ {
			if bytes.Equal(pattern, initial_drug[i:i+pattern_len]) { // (if) pattern found
				result := make([]byte, 0, len(initial_drug)-pattern_len+len(drug.replacement))
				result = append(result, initial_drug[0:i]...)
				result = append(result, drug.replacement...)
				result = append(result, initial_drug[i+pattern_len:]...)
				drugs_set.Append(string(result))
			}
		}
	}
	return len(drugs_set)
}

// ngl my own solution was almost correct but there something wrong with it that i couldn't figure it
// the below solution is from this reddit thread: https://www.reddit.com/r/adventofcode/comments/3xflz8/day_19_solutions/
// the answer is from the user: https://www.reddit.com/user/CdiTheKing/
//Func<string, int> countStr = x =>
//{
//	var count = 0;
//	for (var index = str.IndexOf(x); index >= 0; index = str.IndexOf(x, index + 1), ++count) { }
//	return count;
//};
//
//var num = str.Count(char.IsUpper) - countStr("Rn") - countStr("Ar") - 2 * countStr("Y") - 1;
// the below code is the same code above, translated to go
// P.S my solution gave the answer 190 while the actual solution was 195 :(

func part2(drugs []Drug, initial_drug []byte) int {
	_ = drugs
	str := string(initial_drug)
	countStr := func(substr string) int {
		return strings.Count(str, substr)
	}

	uppercaseCount := 0
	for _, r := range str {
		if r >= 'A' && r <= 'Z' {
			uppercaseCount++
		}
	}

	result := uppercaseCount -
		countStr("Rn") -
		countStr("Ar") -
		2*countStr("Y") -
		1

	return result
}

func Solution1(f *os.File) {
	sc := bufio.NewScanner(f)

	index := 0
	var drugs [43]Drug
	for sc.Scan() {
		line := sc.Bytes()
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

	fmt.Println("the solution to day19 part 1 is:", result)
	f.Seek(0, io.SeekStart)
}

func Solution2(f *os.File) {
	sc := bufio.NewScanner(f)

	index := 0
	var drugs [43]Drug
	for sc.Scan() {
		line := sc.Bytes()
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

	fmt.Println("the solution to day19 part 2 is:", result)
}
