package day24

import (
	"Aoc2015/lib/utility"
	"bufio"
	"fmt"
	"os"
)

func setWeights(in []byte, weights []int, index int) (parsing_error error) {
	val, err := utility.BytesToIntSigned(in)
	if err != nil {
		return err
	}

	weights[index] = val
	return nil
}

var weights = make([]int, 28)

func part1(weights []int) int {
	fmt.Printf("the sum: %d, the groups: %d\n", utility.SumSlice(weights), utility.SumSlice(weights)/3)
	return 0
}

func part2(weights []int) int {
	return 0
}

func Solution1(f *os.File) {
	sc := bufio.NewScanner(f)
	index := 0

	for sc.Scan() {
		line := sc.Bytes()
		_ = line
		if err := setWeights(line, weights, index); err != nil {
			fmt.Println("problem parsing the weights:", err.Error())
			return
		}
		index += 1
	}
	result := part1(weights)

	if err := sc.Err(); err != nil {
		panic(fmt.Sprintf("there was a problem reading the file: %s", err.Error()))
	}

	fmt.Println("the solution to day24 part 1 is:", result)
}

func Solution2(f *os.File) {
	// we don't even need to parse the file again
	result := part2(weights)

	fmt.Println("the solution to day24 part 2 is:", result)
}
