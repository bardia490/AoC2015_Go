package day18

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func generateLightsArray(array []int, in string, row int) {
	columns_count := len(in)

	for index, val := range in {
		value := 0 // if val == '.'
		if val == '#' {
			value = 1
		}
		array[row*columns_count+index] = value
	}
}

func part1(array []int) int {
	result := 0
	_ = array
	return result
}

func part2(array []int) int {
	result := 0
	_ = array
	return result
}

func Solution1(f *os.File) {
	sc := bufio.NewScanner(f)

	index := 0
	var array [100 * 100]int
	for sc.Scan() {
		line := sc.Text()
		generateLightsArray(array[:], line, index)
		index += 1
	}
	for index := 0; index < len(array); index += 100 {
		fmt.Println(index, array[index:index+100])
	}
	result := part1(array[:])

	if err := sc.Err(); err != nil {
		panic(fmt.Sprintf("there was a problem reading the file: %s", err.Error()))
	}

	fmt.Println("the solution to day17 part 1 is:", result)
	f.Seek(0, io.SeekStart)
}

func Solution2(f *os.File) {
	sc := bufio.NewScanner(f)

	index := 0
	var array [1000 * 1000]int
	for sc.Scan() {
		line := sc.Text()
		generateLightsArray(array[:], line, index)
		index += 1
	}
	result := part2(array[:])

	if err := sc.Err(); err != nil {
		panic(fmt.Sprintf("there was a problem reading the file: %s", err.Error()))
	}

	fmt.Println("the solution to day17 part 2 is:", result)
}
