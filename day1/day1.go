package day1

import (
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

func Solution1(in []byte) int {
	result := 0
	contents := strings.SplitSeq(string(in), "\n")

	for line := range contents {
		// process the contents
		result += MoveFloors(line)
	}
	return result
}

type solution2Error struct{}

func (e solution2Error) Error() string {
	return "could not find the solution to the problem"
}

func Solution2(in []byte) (int, error) {
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
				return index + 1, nil
			}
		}
	}
	// if we reached here something has gone wrong and no answer has been found
	return 0, solution2Error{}
}
