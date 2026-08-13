package day1

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func Add(a, b int) int {
	return a + b
}

func MoveFloors(in string) int {
	var result int = 0

	for _, elem := range in {
		switch elem {
		case '(':
			result += 1
		case ')':
			result -= 1
		}
	}

	return result
}

func Solution1(f *os.File) {
	in, _ := io.ReadAll(f)
	f.Seek(0, io.SeekStart)
	result := 0
	contents := strings.SplitSeq(string(in), "\n")

	for line := range contents {
		// process the contents
		result += MoveFloors(line)
	}
	fmt.Println("the solution to day1 part 1 is:", result)
}

type solution2Error struct{}

func (e solution2Error) Error() string {
	return "could not find the solution to the problem"
}

func Solution2(f *os.File) {
	in, _ := io.ReadAll(f)
	result := 0
	contents := strings.SplitSeq(string(in), "\n")

	for line := range contents {
		// process the contents
		for index, character := range line {
			switch character {
			case '(':
				result += 1
			case ')':
				result -= 1
			}
			if result == -1 {
				fmt.Println("the solution to day1 part 2 is:", index+1)
				return
			}
		}
	}
	// if we reached here something has gone wrong and no answer has been found
	fmt.Println("something went wrong with the part2 solution of day1 and the basement was never found")
}
